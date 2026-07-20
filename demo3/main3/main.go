package main3

import "sync"

type User struct {
	Name  string
	Age   int
	Email string
}

type SimpleUser struct {
	Name string
	Age  int
}

type RWMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewRWMap() *RWMap {
	return &RWMap{
		m: make(map[string]int),
	}
}

func (rm *RWMap) Set(key string, value int) {
	rm.mu.Lock() // 写锁
	defer rm.mu.Unlock()
	rm.m[key] = value
}

func (rm *RWMap) Get(key string) (int, bool) {
	rm.mu.RLock() // 读锁，多个读可以并发
	defer rm.mu.RUnlock()
	val, ok := rm.m[key]
	return val, ok
}
