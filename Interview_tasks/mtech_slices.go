package main

import (
	"fmt"
)

// Задача:
// Что выведет код и почему?

func modify(s []int, n int) { // slice [ *[]int, len, cap ]
	s = append(s, n) // [1, 2, 0, 55] len= 4 cap=5  //[1,2,66] - меняется основной массив по ссылке (не лен и капасити в мейн)
	s[0] = 99        // [99, 2, 0, 55] len=4 cap=5  //[99,2,66]
}

func main() {
	s1 := make([]int, 3, 5)  // * int[5] = [0 0 0 0 0]		[0 0 0]
	s2 := s1[:2]             // * int[5] = [0 0 0 0 0]		[0, 0]
	s1[0] = 1                // * int[5] = [1 0 0 0 0]   	[1, 0, 0]
	s2[1] = 2                // * int[5] = [1 2 0 0 0]   	[1, 2] 0, 0, 0
	modify(s1, 55)           //  [99, 2, 0]
	modify(s2, 66)           //
	fmt.Println(cap(s1), s1) // [99 2 66] len=3 cap=5
	fmt.Println(cap(s2), s2) // [99 2] len=2 cap=5
}

// s1: [99 2 66]
// s2: [99 2]
// cap s1: 5 cap s2: 5
