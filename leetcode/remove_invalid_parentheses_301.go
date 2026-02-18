package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	res := make([]string, 0)
	if isValidPar(s) {
		res = append(res, s)
	}
	res = recurseFinding([]rune(s), []rune(s), res, len(s), 1)
	return res
}

func recurseFinding(s []rune, src []rune, res []string, i int, j int) []string {
	// 0   - 4    *123 i=4/j=1
	// 0-2 + 2-4  0*23 i=3/j=1
	// 0-3 + 3-4  01*3 i=2/j=1
	// 0   - len  012* i=1/j=1

	// 0    - 4   **23  i=4/j=2
	// 0-1 + 3-4  0**3  i=3/j=2
	// 0-len	  01**  i=2/j=2

	// 0    - 4   ***3   i=4 / j=3
	// 0   -len	  0***   i=3 / j=3

	if len(s)-j < 1 {
		return res
	}

	if isValidPar(string(src)) {
		res = append(res, string(src))
	}

	//if i == 1 || len(s)-i+j > len(res) {
	//	j++
	//	i = len(s)
	//}
	//i--
	////fmt.Println(string(src), len(src), i, j)
	//src = append(s[0:len(s)-i], s[len(s)-i+j:]...)
	//fmt.Println("src =", string(src))
	return recurseFinding(s, src, res, i, j)
}

func isValidPar(s string) bool {
	count := 0

	for _, v := range s {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
		}

		if count < 0 {
			return false
		}
	}
	return count == 0
}

func main() {
	fmt.Println("res:", removeInvalidParentheses("0123"))

	//fmt.Println("res:", removeInvalidParentheses("()())()"))
	//fmt.Println(removeInvalidParentheses("(a)())()"))
	//fmt.Println(removeInvalidParentheses(")("))

}
