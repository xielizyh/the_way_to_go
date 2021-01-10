package store

import (
	"encoding/gob"
	key "goto_v2/key"
	"io"
	"log"
	"os"
	"sync"
)

// URLStore 存储URL
type URLStore struct {
	urls map[string]string
	// 读写锁允许同一时间多个线程对变量进行读操作，但是只能一个线程进行写操作
	mu   sync.RWMutex
	file *os.File
}

type record struct {
	Key, URL string
}

// NewURLStore 工厂函数
func NewURLStore(filename string) *URLStore {
	s := &URLStore{urls: make(map[string]string)}
	// 0644的前导0表示八进制，644跟linux文件权限表示方法一样
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("URLStore:", err)
	}
	s.file = f
	if err := s.load(); err != nil {
		log.Println("Error loading data in URLStore:", err)
	}
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
			if err := s.save(key, url); err != nil {
				log.Println("Error saving to URLStore:", err)
			}
			return key
		}
	}

	// return ""
}

// save 以gob编码形式存入磁盘
func (s *URLStore) save(key, url string) error {
	// 新建编码器
	e := gob.NewEncoder(s.file)
	// 将record进行编码
	return e.Encode(record{key, url})
}

func (s *URLStore) load() error {
	if _, err := s.file.Seek(0, 0); err != nil {
		return err
	}
	// 新建解码器
	d := gob.NewDecoder(s.file)
	var err error
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
