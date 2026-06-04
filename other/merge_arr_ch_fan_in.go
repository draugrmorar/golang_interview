package main

import (
	"fmt"
	"sync"
)

func MergeChannels(channels []<-chan int) <-chan int {
	res := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range ch {
				res <- c
			}
		}()
	}
	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func main() {
	pr := make([]chan int, 10)
	wg := sync.WaitGroup{}
	for i := range pr {
		pr[i] = make(chan int, 10)
		for y := range 10 {
			wg.Add(1)
			go func(k int) {
				defer wg.Done()
				pr[i] <- (i*10 + k)
			}(y)

		}
		go func() {
			wg.Wait()
			close(pr[i])
		}()
	}

	chans := make([]<-chan int, 10)
	for i := range chans {
		chans[i] = pr[i]
	}
	res := MergeChannels(chans)
	for r := range res {
		fmt.Println(r)
	}
}
