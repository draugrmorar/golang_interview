package main

import (
	"fmt"
	"time"
)

// Задача:
// - что выведет код? (6 сек)
// - сделай так, чтобы вывел 3 сек
func main() {
	timeStart := time.Now()
	_, _ = <-worker(), <-worker()
	fmt.Println(time.Since(timeStart).Seconds())
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}
