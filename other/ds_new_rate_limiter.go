package main

import (
	"sync"
	"time"
)

// Паттерн: Token bucket, select с default
// Ключевое: начальные жетоны для burst, тикер пополняет, select/default отбрасывает лишнее.
func NewRateLimiter2(limit int, fn func()) func() {
	tokens := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		tokens <- struct{}{}
	}

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(limit))
		defer ticker.Stop()
		for range ticker.C {
			select {
			case tokens <- struct{}{}:
			default:
			}
		}
	}()

	var mu sync.Mutex
	return func() {
		select {
		case <-tokens:
			mu.Lock()
			fn()
			mu.Unlock()
		default:
		}
	}
}

// Паттерн: Мьютекс + блокировка на тикере
// Ключевое: мьютекс даёт FIFO, <-ticker.C блокирует до следующего слота.
func NewRateLimiterBlocking(limit int, fn func()) func() {
	ticker := time.NewTicker(time.Second / time.Duration(limit))
	var mu sync.Mutex

	return func() {
		mu.Lock()
		<-ticker.C
		fn()
		mu.Unlock()
	}
}
