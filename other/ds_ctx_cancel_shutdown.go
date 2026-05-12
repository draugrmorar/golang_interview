package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// Условие:
// Есть сервис, который запускает N воркеров. Нужно реализовать graceful shutdown:
// при получении SIGINT или SIGTERM сервис должен:
// - Прекратить приём новых задач
// - Дождаться завершения всех текущих задач (максимум 5 секунд)
// - Вывести лог: сколько задач успешно завершено, сколько отменено
// Воркеры могут выполнять долгие операции, которые нужно уметь прерывать через context.

type workerPool struct {
	numOfWorkers int
	workCh       chan Task

	wg sync.WaitGroup

	ctx    context.Context
	cancel context.CancelFunc

	numDone   int64
	numCancel int64
}

type Task struct {
	ID int
	fn func(ctx context.Context) error
}

func newWorkerPool(numOfWorkers int, workNum int) *workerPool {
	ctx, cancel := context.WithCancel(context.Background())

	return &workerPool{
		numOfWorkers: numOfWorkers,
		workCh:       make(chan Task, workNum),
		wg:           sync.WaitGroup{},
		ctx:          ctx,
		cancel:       cancel,
	}
}

func (p *workerPool) start() {
	for i := 0; i < p.numOfWorkers; i++ {
		p.wg.Add(1)
		go p.work(i)
	}
}

func (p *workerPool) work(wnum int) {
	defer p.wg.Done()
	for {
		select {
		case task, ok := <-p.workCh:
			if !ok {
				return
			}
			err := task.fn(p.ctx)
			if err != nil {
				atomic.AddInt64(&p.numCancel, 1)
				fmt.Printf("worker %d failed: %v\n", wnum, err)
			} else {
				atomic.AddInt64(&p.numDone, 1)
				fmt.Printf("worker %d done\n", wnum)
			}
		case <-p.ctx.Done():
			return
		}
	}
}

func (p *workerPool) addTask(task Task) bool {
	select {
	case p.workCh <- task:
		return true
	case <-p.ctx.Done():
		return false
	}
}

func (p *workerPool) Stop(timeout time.Duration) (completed, canceled int64) {
	p.cancel()
	close(p.workCh)

	done := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		fmt.Println("All workers done")
	case <-time.After(timeout):
		fmt.Println("All workers timeout")
	}
	return atomic.LoadInt64(&p.numDone), atomic.LoadInt64(&p.numCancel)
}

func main() {
	work := []string{"http://www.baidu.com",
		"https://www.ya.ru",
		"http://google.com",
		"ghjk",
		"http://youtube.com",
		"http://google.com",
		"http://twitter.com",
		"http://bing.com",
	}
	pool := newWorkerPool(2, len(work))
	pool.start()
	go func() {
		for i, w := range work {
			task := Task{
				ID: i,
				fn: func(ctx context.Context) error {
					req, err := http.NewRequestWithContext(ctx, "GET", w, nil)
					if err != nil {
						return err
					}
					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						return err
					}
					defer resp.Body.Close()
					return nil
				},
			}
			if !pool.addTask(task) {
				fmt.Printf("Task %d rejected\n", i)
				break
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	// Graceful shutdown
	fmt.Println("\nShutting down...")
	completed, canceled := pool.Stop(5 * time.Second)
	fmt.Printf("Completed: %d, Canceled: %d\n", completed, canceled)
}
