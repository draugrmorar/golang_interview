package main

import "sync"

// с wait group
func workerPool(jobs []int, workers int) {
	var wg sync.WaitGroup
	jobCh := make(chan int, len(jobs))
	// Запускаем workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobCh {
				process(job) // обрабатываем задачу
			}
		}()
	}
	// Отправляем задачи
	for _, job := range jobs {
		jobCh <- job
	}
	close(jobCh) // закрываем, чтобы workers завершились
	wg.Wait()    // ждем всех workers
}

// с семафором
func processWithLimit(jobs []Job, maxWorkers int) {
	sem := NewSemaphore(maxWorkers)
	var wg sync.WaitGroup

	for _, job := range jobs {
		wg.Add(1)
		go func(j Job) {
			defer wg.Done()

			sem.Acquire() // ждем свободного worker'а
			defer sem.Release()

			process(j) // делаем работу
		}(job)
	}

	wg.Wait()
}
