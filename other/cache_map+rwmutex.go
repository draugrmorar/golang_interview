package main

import "sync"

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
}

func (c *Cache) Get(key string) string {
	c.mu.RLock()         // блокировка ЧТЕНИЯ
	defer c.mu.RUnlock() // обязательно!

	return c.data[key]
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock() // эксклюзивная блокировка
	defer c.mu.Unlock()

	c.data[key] = value
}

func (c *Cache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.data)
}
