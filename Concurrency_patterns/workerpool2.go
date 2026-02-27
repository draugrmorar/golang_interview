package main

import (
	"fmt"
	"time"
)

func works(jobsCh <-chan string, rCh chan<- string) {
	for jCh := range jobsCh {
		fmt.Printf("input: %s\n", jCh)
		time.Sleep(time.Second)
		rCh <- jCh
	}
}

func main() {
	wn := 3
	jobs := []string{"1job", "2job", "3job", "4job", "5job", "6job", "7job", "8job", "9job"}
	resCh := make(chan string, len(jobs))
	jobsCh := make(chan string, len(jobs))

	for range wn {
		go works(jobsCh, resCh)
	}
	for _, j := range jobs {
		jobsCh <- j
	}
	close(jobsCh)
	for range len(jobs) {
		fmt.Printf("res %s\n", <-resCh)
	}
}
