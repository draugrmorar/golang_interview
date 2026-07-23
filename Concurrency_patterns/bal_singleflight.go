package main

import (
	"fmt"
	"sync"
	"time"
)

type call struct {
	value interface{}
	err   error
	sigCh chan struct{}
}

type SingleFlight struct {
	mx    sync.Mutex
	calls map[string]*call
}

func NewSingleFlight() *SingleFlight {
	return &SingleFlight{
		mx:    sync.Mutex{},
		calls: make(map[string]*call),
	}
}

func (s *SingleFlight) Do(key string, action func() (interface{}, error)) (interface{}, error) {
	s.mx.Lock()
	if call, ok := s.calls[key]; ok {
		s.mx.Unlock()
		return s.Wait(call)
	}

	call := &call{
		sigCh: make(chan struct{}),
	}
	s.calls[key] = call
	s.mx.Unlock()

	go func() {
		defer func() {
			s.mx.Lock()
			close(call.sigCh)
			//delete(s.calls, key)
			s.mx.Unlock()
		}()
		s.calls[key].value, s.calls[key].err = action()
	}()

	return s.Wait(call)
}

func (s *SingleFlight) Wait(call *call) (interface{}, error) {
	<-call.sigCh
	return call.value, call.err
}

func main() {
	const inFlightRequests = 5
	wg := sync.WaitGroup{}
	wg.Add(inFlightRequests)

	singleFlight := NewSingleFlight()

	const key = "same_key"
	for i := 0; i < inFlightRequests; i++ {
		go func() {
			defer wg.Done()
			value, err := singleFlight.Do(key, func() (interface{}, error) {
				fmt.Println("single flight")
				time.Sleep(5 * time.Second)
				return "result", nil
			})

			fmt.Println(i, "=", value, err)
		}()
	}

	wg.Wait()
}
