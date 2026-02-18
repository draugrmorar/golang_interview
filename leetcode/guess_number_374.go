package main

import "fmt"

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 *
 */

func guessNumber(n int) int {
	var g int
	var mn int
	var num int
	for mn <= n {
		num = (n + mn) / 2
		g = guess(num)
		if g == -1 {
			n = num - 1
		} else if g == 1 {
			mn = num + 1
		} else {
			return num
		}
	}
	return n
}

func guess(num int) int {
	pick := 1
	if num == pick {
		return 0
	} else if num > pick {
		return -1
	}
	return 1
}

func main() {
	fmt.Println(guessNumber(1))
}
