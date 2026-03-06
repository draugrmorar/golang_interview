package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"http://www.baidu.com",
		"https://www.ya.ru",
		"http://google.com",
		"ghjk",
	}
	resCh := make(chan string, len(urls))
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				resCh <- "resp is not OK: " + err.Error()
				return
			}
			defer resp.Body.Close()
			resCh <- "resp status code for url " + url + " is " + resp.Status
		}()
	}
	wg.Wait()
	close(resCh)

	for resp := range resCh {
		fmt.Println(resp)
	}
}
