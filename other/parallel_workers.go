package main

import (
	"fmt"
	"sync"
)

func ParallelWorkers(tasks <-chan int, fn func(int) int, workers int) []int {
	var res []int
	wg := sync.WaitGroup{}
	resCh := make(chan int, workers)
	for range workers {
		wg.Add(1)
		go func(taskss <-chan int) {
			defer wg.Done()
			for t := range taskss {
				resCh <- fn(t)
			}
		}(tasks)
	}
	go func() {
		wg.Wait()
		close(resCh)
	}()
	for r := range resCh {
		res = append(res, r)
	}
	return res
}

func Multi(i int) int {
	return i * i
}

func main() {
	in := []int{2, 3, 4, 5, 6, 8, 9, 2, 3, 4, 5, 6, 8, 9}
	ch := make(chan int, len(in))
	for i := range in {
		ch <- in[i]
	}
	close(ch)
	fmt.Println(ParallelWorkers(ch, Multi, 3))
}
