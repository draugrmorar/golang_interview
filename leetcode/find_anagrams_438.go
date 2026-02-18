package main

import (
	"fmt"
	"slices"
)

func findAnagrams(s string, p string) []int {
	r1 := []rune(p)
	slices.Sort(r1)
	l := len(p)
	res := make([]int, 0, len(s))
	for i := 0; i <= len(s)-l; i++ {
		r2 := []rune(s[i : i+l])
		slices.Sort(r2)
		if slices.Equal(r1, r2) {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
}
