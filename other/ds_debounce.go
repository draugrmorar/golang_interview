package main

import (
	"sync"
	"time"
)

// Паттерн: time.AfterFunc + timer.Stop
// Ключевое: каждый новый вызов сбрасывает таймер, fn выполняется когда вызовы прекратились.
func Debounce(fn func(), d time.Duration) func() {
	var mu sync.Mutex
	var timer *time.Timer

	return func() {
		mu.Lock()
		defer mu.Unlock()
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(d, fn)
	}
}
