package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ErrGroup struct {
	err    error
	wg     sync.WaitGroup
	once   sync.Once
	doneCh chan struct{}
}

func NewErrGroup() *ErrGroup {
	return &ErrGroup{
		doneCh: make(chan struct{}),
	}
}

func (e *ErrGroup) Do(task func() error) {
	select {
	case <-e.doneCh:
		return
	default:
	}
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		select {
		case <-e.doneCh:
			return
		default:
			if err := task(); err != nil {
				e.once.Do(func() {
					e.err = err
					close(e.doneCh)
				})
			}
		}
	}()
}

func (e *ErrGroup) Wait() error {
	e.wg.Wait()
	return e.err
}

func main() {
	errGroup := NewErrGroup()
	errGroup.Do(func() error {
		fmt.Println("start")
		return errors.New("!!! ERROR !!! ")
		//return nil
	})
	time.Sleep(time.Second)
	for i := range 5 {
		errGroup.Do(func() error {
			fmt.Println("started after timeout", i)
			return nil
		})
	}
	if err := errGroup.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}
