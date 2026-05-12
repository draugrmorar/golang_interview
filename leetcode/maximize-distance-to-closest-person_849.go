package main

import "fmt"

func maxDistToClosest(seats []int) int {
	res, j := 0, 0
	first := true
	for i := range seats {
		if first && seats[i] == 1 {
			res = j
			first = false
			j = 0
		} else if seats[i] == 1 {
			res = max((j+1)/2, res)
			j = 0
		} else {
			j++
		}
	}
	return max(j, res)
}

func main() {
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1})) // 2
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0}))          // 3
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1})) // 2
	fmt.Println(maxDistToClosest([]int{1, 0, 0, 1}))          // 1
}
