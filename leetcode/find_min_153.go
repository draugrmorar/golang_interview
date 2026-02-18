package main

import "fmt"

func findMin(nums []int) int {
	l := len(nums)
	res := nums[l-1]
	if nums[l-1] >= nums[0] {
		return nums[0]
	} else {
		for i := 2; i < l; i++ {
			if res > nums[l-i] {
				res = nums[l-i]
			} else {
				return res
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}
