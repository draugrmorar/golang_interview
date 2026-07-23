package main

import (
	"sync"
	"time"
)

// Паттерн: time.Since + lastCall
// Ключевое: проверка time.Since(lastCall), вызов только если прошло больше d.
func Throttle(fn func(), d time.Duration) func() {
	var lastCall time.Time
	var mu sync.Mutex

	return func() {
		mu.Lock()
		if time.Since(lastCall) >= d {
			lastCall = time.Now()
			mu.Unlock()
			fn()
		} else {
			mu.Unlock()
		}
	}
}
