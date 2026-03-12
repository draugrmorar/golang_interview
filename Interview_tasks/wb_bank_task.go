package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(methodAfterChange([]string{"Petya", "Vasya", "Vanya", "Igor", "Petya"}))
	fmt.Println("Time spent:", time.Since(start))
}

func someWork(s string) (string, error) {
	// work
	time.Sleep(10 * time.Millisecond)
	return s, nil //errors.New("big real error")
}

func methodBeforeChange(str []string) map[string]int {
	nameMap := make(map[string]int)
	for _, s := range str {
		name, err := someWork(s)
		if err != nil {
			fmt.Println(err) // не помню как обрабатывали ошибку, может прокидывали выше
			return nil
		}
		nameMap[name] = nameMap[name] + 1
	}
	return nameMap
}

func methodAfterChange(str []string) map[string]int {
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	errCh := make(chan error, len(str))

	nameMap := make(map[string]int)
	for _, s := range str {
		wg.Add(1)
		go func(y string) {
			defer wg.Done()
			name, err := someWork(y)
			if err != nil {
				errCh <- err
			}
			mx.Lock()
			nameMap[name] = nameMap[name] + 1
			mx.Unlock()
		}(s)
	}
	go func() {
		wg.Wait()
		close(errCh)
	}()
	if err, ok := <-errCh; err != nil && ok {
		fmt.Println(err)
		return nil
	}
	return nameMap
}

// Задача примерно такая была, пишу по памяти
// Что еще спрашивали:
// - мапы / -каналы буфферизированные/ не буфферизированные
// - SQL - что делать если развернут сервер на трех подах кубера и
//   они все делают одни и те же запросы в одну и ту же таблицу
//   какие уровни изоляции есть в бд, от чего защищают
//   select for update
// 	 шардирование как можно решить проблему шардированием
//   как можно решить проблему паттернами? - вот тут не помню про какой паттерн он сказал
// - Kafka - зачем нужны паритиции
//   какие гарантии доставки есть
//   как добиться Exactly-once - какие настройки нужно сделать
//   паттерны для ассинхронного взаимодействия, в том числе transactional outbox/ inbox/  не помню тут тоже какой-то паттерн он назвал
