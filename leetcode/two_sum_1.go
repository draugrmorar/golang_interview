package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var res int
	for i, n := range nums {
		res = target - n
		if v, ok := m[res]; ok {
			return []int{i, v}
		}
		m[n] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
