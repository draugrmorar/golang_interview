package main

import (
	"fmt"
	"regexp"
)

// Example 1:
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:
// Input: s = "cbbd"
// Output: "bb"

// Constraints:
// 1 <= s.length <= 1000
// s consist of only digits and English letters.

func main() {
	fmt.Println(longestPalindrome("bdsdbbb"))
	//  fmt.Println(longestPalindrome("cbbd"))
	//  fmt.Println(longestPalindrome("ccc"))
	//  fmt.Println(longestPalindrome("aa"))
	//  fmt.Println(longestPalindrome("a"))
	//  fmt.Println(longestPalindrome("dfghnmnmnm"))
}

// по i мы перемещаемся по строке
// по y мы отсчитываем вправо и влево
func longestPalindrome(s string) string {
	var res string
	var x int

	for i := 0; i < len(s); i++ {
		x = 0
		// надо посмотреть и 0 значения! типа ссbd
		// шагаем вправо и влево после нахождения одинаковых букв!
		// а то мы шагаем в том же месте в котором находим одинаковые буквы
		for y := 0; (i + y + x) < len(s); y++ {
			fmt.Println(i, y, x, res)
			if y == 0 && i+1 < len(s) && s[i] == s[i+1] {
				if len(res) <= ((y * 2) + x) {
					res = s[i : i+2]
				}
				x = 1
			}
			if !(y == 0 && x == 1) && (i-y) >= 0 {
				fmt.Println(i, y, res)
				if s[i-y] == s[i+y+x] && len(res) <= ((y*2)+x) {
					res = s[i-y : i+y+x+1]
					fmt.Println(res)
				}
			} else {
				break
			}
		}
	}
	return res
}

func parse(str string, ch chan bool) {
	if len(str) < 1 || len(str) > 1000 {
		ch <- false
		return
	}
	nok, _ := regexp.MatchString("^[A-Za-z0-9]+$", str)
	if !nok {
		ch <- false
		return
	}
	ch <- true
}
