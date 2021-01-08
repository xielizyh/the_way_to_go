package store

import (
	key "goto_v1/key"
	"sync"
)

// URLStore 存储URL
type URLStore struct {
	urls map[string]string
	// 读写锁允许同一时间多个线程对变量进行读操作，但是只能一个线程进行写操作
	mu sync.RWMutex
}

// NewURLStore 工厂函数
func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]string)}
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

// Put 接收长URL，计算短URL，存储长URL
func (s *URLStore) Put(url string) string {
	for {
		key := key.GenKey(s.Count())
		if ok := s.Set(key, url); ok {
			return key
		}
	}

	// return ""
}
