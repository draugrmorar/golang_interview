package main

import "fmt"

func search(nums []int, target int) int {
	l := len(nums)
	for i := 0; i < l-i; i++ {
		if nums[i] > target && nums[l-i-1] < target {
			return -1
		}
		if nums[i] == target {
			return i
		}
		if nums[l-i-1] == target {
			return l - i - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))    //4
	fmt.Println(search([]int{4, 5, 6, 8, 9, 0, 1, 2}, 8)) //3
	fmt.Println(search([]int{1}, 0))                      //-1

}
