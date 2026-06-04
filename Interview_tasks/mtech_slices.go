package main

import "fmt"

// Задача:
// Что выведет код и почему?

func modify(s []int, n int) {
	s = append(s, n)
	s[0] = 99
}

func main() {
	s1 := make([]int, 3, 5)
	s2 := s1[:2]
	s1[0] = 1
	s2[1] = 2
	modify(s1, 55)
	modify(s2, 66)
	fmt.Println(cap(s1), s1)
	fmt.Println(cap(s2), s2)
}

// s1: [99 2 66]
// s2: [99 2]
// cap s1: 5 cap s2: 5
