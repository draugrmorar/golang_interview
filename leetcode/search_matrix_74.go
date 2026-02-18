package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	var i int
	if matrix[0][0] > target {
		return false
	}
	for ; i < len(matrix); i++ {
		if matrix[i][0] > target {
			break
		}
	}
	i--
	for j := 0; j < len(matrix[i]) && target >= matrix[i][j]; j++ {
		if matrix[i][j] == target {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1}}, 0))
	fmt.Println(searchMatrix([][]int{{1}}, 2))
}
