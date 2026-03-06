package main

import "fmt"

// Функция, которая принимает два канала, в которые будет поступать все в отсортированном порядке
// Задача функции смержить два отсортированных канала в один и вернуть его.
// Вернуть канал, когда все переданные каналы закроются

func mergeSortedChannels(ch1, ch2 chan int) chan int {
	out := make(chan int)
	s1 := make([]int, 0)
	s2 := make([]int, 0)

	for ch := range ch1 {
		s1 = append(s1, ch)
	}
	for ch := range ch2 {
		s2 = append(s2, ch)
	}
	go sortSlices(s1, s2, out)
	return out
}

func sortSlices(s1, s2 []int, out chan int) {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			out <- s1[i]
			i++
		} else {
			out <- s2[j]
			j++
		}
	}

	for i < len(s1) {
		out <- s1[i]
		i++
	}
	for j < len(s2) {
		out <- s2[j]
		j++
	}
	close(out)
}

// если не отсортированные слайсы
func mergeNotSortedChannels(ch1, ch2 chan int) chan int {
	out := make(chan int)
	res := make([]int, 0)
	for ch := range ch1 {
		res = append(res, ch)
	}
	for ch := range ch2 {
		res = append(res, ch)
	}
	go sortNotSortedSlice(res, out)
	return out
}

func sortNotSortedSlice(s []int, out chan int) {
	s = bubbleSort(s)
	for i := range s {
		out <- s[i]
	}
	close(out)
}

func bubbleSort(s []int) []int {
	for i := 0; i+1 < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func fillCh1(ch1 chan int) {
	ch1 <- 1
	ch1 <- 3
	ch1 <- 4
	close(ch1)
}

func fillCh2(ch2 chan int) {
	ch2 <- 3
	ch2 <- 5
	ch2 <- 10
	ch2 <- 15
	close(ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go fillCh1(ch1)
	go fillCh2(ch2)
	for num := range mergeSortedChannels(ch1, ch2) {
		fmt.Println(num)
	}
}
