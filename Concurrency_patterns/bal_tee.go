package main

import (
	"fmt"
	"sync"
)

func tee(in chan int) (chan int, chan int) {
	out1 := make(chan int)
	out2 := make(chan int)
	go func() {
		defer close(out1)
		defer close(out2)
		for i := range in {
			out1 <- i
			out2 <- i
		}
	}()
	return out1, out2
}

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			ch <- i
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch1, ch2 := tee(ch)
	go func() {
		defer wg.Done()
		for r := range ch1 {
			fmt.Println("1: ", r)
		}
	}()
	go func() {
		defer wg.Done()
		for r := range ch2 {
			fmt.Println("2: ", r)
		}
	}()
	wg.Wait()
}
