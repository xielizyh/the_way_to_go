package store

import (
	"encoding/gob"
	"errors"
	key "goto_v5/key"
	"io"
	"log"
	"net/rpc"
	"os"
	"sync"
)

const saveQueueLength = 1024

// Store 存储接口
type Store interface {
	Put(key, url *string) error
	Get(key, url *string) error
}

// ProxyStore 代理缓存
type ProxyStore struct {
	urls   *URLStore
	client *rpc.Client
}

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
	s := &URLStore{urls: make(map[string]string)}
	if filename != "" {
		s.save = make(chan record, saveQueueLength)
		if err := s.load(filename); err != nil {
			log.Println("Error loading data in URLStore:", err)
		}
		go s.saveLoop(filename)
	}

	return s
}

// Get 获取
func (s *URLStore) Get(key, url *string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if u, ok := s.urls[*key]; ok {
		*url = u
		return nil
	}
	return errors.New("key not founc")
}

// Set 设置
func (s *URLStore) Set(key, url *string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	// 如果键存在，map不进行更新
	if _, present := s.urls[*key]; present {
		return errors.New("key already exists")
	}
	s.urls[*key] = *url
	return nil
}

// dount 获取键值对的数量
func (s *URLStore) count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}

// Put 接收长URL，计算短URL，存储数据
func (s *URLStore) Put(url, k *string) error {
	// fmt.Printf("%s:%s", *url, *k)
	for {
		*k = key.GenKey(s.count())
		if err := s.Set(k, url); err == nil {
			// 发送记录record到信道
			s.save <- record{*k, *url}
			return nil
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
	// d := json.NewDecoder(f)
	for err == nil {
		var r record
		// 解码
		if err = d.Decode(&r); err == nil {
			s.Set(&r.Key, &r.URL)
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
	// e := json.NewEncoder(f)
	for {
		r := <-s.save
		if err := e.Encode(r); err != nil {
			log.Println("Error saving to URLStore: ", err)
		}
	}
}

// NewProxyStore 创建ProxyStore
func NewProxyStore(addr string) *ProxyStore {
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Println("Error constructing ProxyStore: ", err)
	}
	return &ProxyStore{urls: NewURLStore(""), client: client}
}

// Get 代理存储Get
func (s *ProxyStore) Get(key, url *string) error {
	// url found in local map
	if err := s.urls.Get(key, url); err == nil {
		return nil
	}
	// url not found in local map, make rpc-call:
	if err := s.client.Call("Store1.Get", key, url); err != nil {
		return err
	}
	s.urls.Set(key, url)
	return nil
}

// Put 代理存储Put
func (s *ProxyStore) Put(url, key *string) error {
	// 注意这里远程调用的第一个参数：Type.Method
	// 其中Type必须为rpc.RegisterName注册的名字(如果使用rpc.Register默认就是对象的类型名字)
	if err := s.client.Call("Store1.Put", url, key); err != nil {
		// if err := s.client.Call("st.Put", url, key); err != nil {
		return err
	}
	s.urls.Set(key, url)
	return nil
}
