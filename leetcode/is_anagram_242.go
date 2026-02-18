package main

import (
	"fmt"
	"slices"
	"sort"
)

func isAnagram(str string, str2 string) bool {
	if len(str) != len(str2) {
		return false
	}
	r := []rune(str)
	r2 := []rune(str2)
	slices.Sort(r)
	slices.Sort(r2)
	for i := range r {
		if r[i] != r2[i] {
			return false
		}
	}
	return true
}

func isAnagram2(str string, str2 string) bool {
	if len(str) != len(str2) {
		return false
	}
	nums1 := make([]int, len(str))
	nums2 := make([]int, len(str2))
	for i := range str {
		nums1[i] = int(str[i])
		nums2[i] = int(str2[i])
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i := range nums1 {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
