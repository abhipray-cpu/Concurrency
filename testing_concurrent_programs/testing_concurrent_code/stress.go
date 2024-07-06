package main

import (
	"sync"
)

type InMemoryDB struct {
	data map[string]string
	lock sync.RWMutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]string),
	}
}

func (db *InMemoryDB) Set(key, value string) {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.data[key] = value
}

func (db *InMemoryDB) Get(key string) (string, bool) {
	db.lock.RLock()
	defer db.lock.RUnlock()
	value, exists := db.data[key]
	return value, exists
}
