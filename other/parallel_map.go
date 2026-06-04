package main

import (
	"fmt"
	"sync"
)

func ParallelMap(input []int, fn func(int) int, concurrency int) []int {
	res := []int{}
	jobCh := make(chan int, len(input))
	wg := sync.WaitGroup{}
	for range concurrency {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for o := range jobCh {
				res = append(res, fn(input[o]))
			}
		}()
	}
	for i := 0; i < len(input); i++ {
		jobCh <- i
	}
	close(jobCh)
	wg.Wait()
	return res
}

func main() {
	fmt.Println(ParallelMap([]int{2, 3, 4, 5, 6, 8, 9, 2, 3, 4, 5, 6, 8, 9}, increment, 3))
	fmt.Println(ParallelMap([]int{2, 3, 4, 5, 6, 8, 9, 2, 3, 4, 5, 6, 8, 9}, multi, 3))

}

func increment(i int) int {
	return i + 1
}
func multi(i int) int {
	return i * 2
}
