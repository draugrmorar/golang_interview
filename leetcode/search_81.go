package main

import "fmt"

func search(nums []int, target int) bool {
	for i := 0; i < len(nums)-i; i++ {
		if nums[i] == target || nums[len(nums)-i-1] == target {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{4, 5, 6, 0, 0, 1, 4}, 3))
	fmt.Println(search([]int{2, 2, 2, 3, 2, 2, 2}, 3))
}
