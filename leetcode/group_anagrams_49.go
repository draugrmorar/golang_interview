package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	res = append(res, []string{strs[0]})
	var isAn bool
	for j := 1; j < len(strs); j++ {
		isAn = false
		for i := 0; i < len(res); i++ {
			if isAnagram(strs[j], res[i][0]) {
				res[i] = append(res[i], strs[j])
				isAn = true
				break
			}
		}
		if !isAn {
			res = append(res, []string{strs[j]})
		}
	}
	return res
}

// Задача 242
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

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "nat", "bat"}))
}
