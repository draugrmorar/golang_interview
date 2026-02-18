package main

import "fmt"

func singleNumber(nums []int) int {
	var sum int
	for i := 0; i+1 < len(nums); i++ {
		sum += nums[i]
		sum -= nums[i+1]
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 2, 3, 4, 5}
	fmt.Println(singleNumber(nums))
}
