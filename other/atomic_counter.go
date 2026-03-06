package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var res int64
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			atomic.AddInt64(&res, int64(i))
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&res))
}
