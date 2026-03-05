package main

import (
	"fmt"
	"time"
)

func worker2(jCh <-chan string, resCh chan<- string) {
	for j := range jCh {
		fmt.Printf("j started: %s\n", j)
		time.Sleep(3 * time.Second)
		resCh <- j
	}
}

func main() {
	jobs := []string{"a", "b", "c", "d", "e", "j", "k", "l"}
	jCh := make(chan string, len(jobs))
	resCh := make(chan string, len(jobs))
	for range 3 {
		go worker2(jCh, resCh)
	}
	for j := range len(jobs) {
		jCh <- jobs[j]
	}
	close(jCh)
	for range len(jobs) {
		fmt.Printf("Res: %s\n", <-resCh)
	}
}

//func main() {
//	var wg sync.WaitGroup
//	s := NewSema(3)
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(j int) {
//			defer wg.Done()
//			s.Aquire()
//			defer s.Release()
//			fmt.Printf("work: %d\n", j)
//			time.Sleep(3 * time.Second)
//		}(i)
//	}
//	wg.Wait()
//}
//
//type Sem struct {
//	semaCh chan struct{}
//}
//
//func NewSema(N int) *Sem {
//	return &Sem{make(chan struct{}, N)}
//}
//
//func (s *Sem) Aquire() {
//	s.semaCh <- struct{}{}
//}
//
//func (s *Sem) Release() {
//	<-s.semaCh
//}

//func main() {
//	var mx sync.WaitGroup
//	s := NewSema(2)
//	for i := 0; i < 10; i++ {
//		mx.Add(1)
//		go func(j int) {
//			s.Acquire()
//			defer mx.Done()
//			defer s.Release()
//			fmt.Printf("%d\n", j)
//			time.Sleep(3 * time.Second)
//		}(i)
//	}
//	mx.Wait()
//}
//
//type Sema struct {
//	sCh chan struct{}
//}
//
//func NewSema(n int) *Sema {
//	return &Sema{
//		sCh: make(chan struct{}, n),
//	}
//}
//
//func (s *Sema) Acquire() {
//	s.sCh <- struct{}{}
//}
//
//func (s *Sema) Release() {
//	<-s.sCh
//}
