package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Реализовать middleware для HTTP хендлера, который ограничивает количество запросов до 10 в секунду с одного IP-адреса.
// Если лимит превышен — возвращать 429 Too Many Requests.
// Нужно потокобезопасное решение

type RateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	timestamps := rl.requests[ip]

	valid := make([]time.Time, 0)
	for _, ts := range timestamps {
		if now.Sub(ts) <= rl.window {
			valid = append(valid, ts)
		}
	}

	if len(valid) >= rl.limit {
		return false
	}

	valid = append(valid, now)
	rl.requests[ip] = valid
	return true
}

func RateLimitMiddleware(rl *RateLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.Header.Get("X-Real-IP")
			if !rl.Allow(ip) {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("Too Many Requests\n"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	limiter := NewRateLimiter(2, 10*time.Second)
	http.Handle("/", RateLimitMiddleware(limiter)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello!\n"))
		if err != nil {
			return
		}
	})))
	fmt.Println("Server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
