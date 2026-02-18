package main

import "fmt"

func search(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5}, 5))
}
