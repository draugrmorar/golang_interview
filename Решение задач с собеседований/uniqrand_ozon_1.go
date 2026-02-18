package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(uniqRandn(20))
}

// Функция должна возвращать n рандомных уникальных чисел
func uniqRandn(n int) []int {
	res := make([]int, 0, n)
	m := make(map[int]struct{})
	for i := 0; i < n; {
		a := rand.Intn(100)
		if _, ok := m[a]; !ok {
			m[a] = struct{}{}
			i++
		}
	}
	for k := range m {
		res = append(res, k)
	}
	return res
}
