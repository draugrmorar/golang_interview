package main

import (
	"context"
	"fmt"
	"time"
)

// 1. WithCancel
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(1 * time.Second) // работаем
	cancel()                    // отменяем
	time.Sleep(1 * time.Second) // время на завершение
}

// 2. WithTimeout
// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // спустя 2 секунды воркер перестанет работать

// 3. WithDeadline. Можно указать точное время остановки.
// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second)) // прибавляем к текущему времени две секунды
