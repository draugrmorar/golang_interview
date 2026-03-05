package main

import (
	"fmt"
)

func main() {
	list := []int{1, 3, 4, 5}
	q := func(x int, y int) int {
		return x * y
	}
	mx := func(x int, y int) bool {
		return x > y
	}
	fmt.Println(Map(list, q, 2))
	fmt.Println(Filter(list, mx, 3))
}

func Map[T any](list []T, f func(T, T) T, opt T) []T {
	res := make([]T, len(list))
	for i, v := range list {
		res[i] = f(v, opt)
	}
	return res
}

func Filter[T any](list []T, f func(T, T) bool, opt T) []T {
	res := make([]T, 0, len(list))
	for _, v := range list {
		if f(v, opt) {
			res = append(res, v)
		}
	}
	return res
}
