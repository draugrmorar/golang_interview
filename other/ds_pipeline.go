package main

import "sync"

// Паттерн: Две WaitGroup, каналы между этапами, запись по индексу
// Ключевое: две WaitGroup, close(midCh) после завершения этапа 1, асинхронный конвейер.
func Pipeline(input []int) []int {
	n := len(input)
	res := make([]int, n)
	idxCh := make(chan int, n)
	midCh := make(chan int, n)
	var wg1, wg2 sync.WaitGroup

	// Этап 1: умножение (3 воркера)
	for range 3 {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			for i := range idxCh {
				res[i] = input[i] * 2
				midCh <- i
			}
		}()
	}

	// Этап 2: квадрат (2 воркера)
	for range 2 {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for i := range midCh {
				res[i] = res[i] * res[i]
			}
		}()
	}

	for i := 0; i < n; i++ {
		idxCh <- i
	}
	close(idxCh)

	wg1.Wait()
	close(midCh)

	wg2.Wait()
	return res
}
