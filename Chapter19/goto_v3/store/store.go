package store

import (
	"encoding/gob"
	key "goto_v3/key"
	"io"
	"log"
	"os"
	"sync"
)

const saveQueueLength = 1024

// URLStore 存储URL
type URLStore struct {
	urls map[string]string
	// 读写锁允许同一时间多个线程对变量进行读操作，但是只能一个线程进行写操作
	mu   sync.RWMutex
	save chan record
}

type record struct {
	Key, URL string
}

// NewURLStore 工厂函数
func NewURLStore(filename string) *URLStore {
	s := &URLStore{
		urls: make(map[string]string),
		save: make(chan record, saveQueueLength),
	}
	if err := s.load(filename); err != nil {
		log.Println("Error loading data in URLStore:", err)
	}
	go s.saveLoop(filename)
	return s
}

// Get 获取
func (s *URLStore) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[key]
}

// Set 设置
func (s *URLStore) Set(key, url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	// 如果键存在，map不进行更新
	if _, present := s.urls[key]; present {
		return false
	}
	s.urls[key] = url
	return true
}

// Count 获取键值对的数量
func (s *URLStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

// Put 接收长URL，计算短URL，存储数据
func (s *URLStore) Put(url string) string {
	for {
		key := key.GenKey(s.Count())
		if ok := s.Set(key, url); ok {
			// 发送记录record到信道
			s.save <- record{key, url}
			return key
		}
	}

	// return ""
}

func (s *URLStore) load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("Error openning URLStore:", err)
		return err
	}
	defer f.Close()
	// 新建解码器
	d := gob.NewDecoder(f)
	for err == nil {
		var r record
		// 解码
		if err = d.Decode(&r); err == nil {
			s.Set(r.Key, r.URL)
		}
	}
	// 解码到最后一条记录
	if err == io.EOF {
		return nil
	}
	return err
}

func (s *URLStore) saveLoop(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening URLStore: ", err)
	}
	defer f.Close()
	e := gob.NewEncoder(f)
	for {
		r := <-s.save
		if err := e.Encode(r); err != nil {
			log.Println("Error saving to URLStore: ", err)
		}
	}
}
