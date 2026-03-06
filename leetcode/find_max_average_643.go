package main

import (
	"fmt"
	"math"
)

//Алгоритм скользящего окна

// Time limit exceeded
func findMaxAverage2(nums []int, k int) float64 {
	maxAv := -math.MaxFloat64
	var midAv float64
	var sum float64
	for i := 0; i+k <= len(nums); i++ {
		for j := 0; j < k; j++ {
			sum += float64(nums[i+j])
		}
		midAv = sum / float64(k)
		if maxAv < midAv {
			maxAv = midAv
		}
		sum = 0
	}
	return maxAv
}

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	for i := 0; i < k; i++ { // сумма самого первого окна
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{-5}, 1))
}
