package main

import "fmt"

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	quickSortRecursive(nums, 0, len(nums)-1)
	return nums
}

func quickSortRecursive(nums []int, low, high int) {
	if low < high {
		pivotIndex := partition(nums, low, high)
		quickSortRecursive(nums, low, pivotIndex-1)
		quickSortRecursive(nums, pivotIndex+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func main() {
	fmt.Println(quickSort([]int{2, 0, 2, 1, 1, 0}))
}
