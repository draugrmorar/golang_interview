package main

import "fmt"

func isValid(s string) bool {
	st := make([]int32, 0, len(s))
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			st = append(st, v)
		} else if len(st) >= 1 && ((v == ')' && st[len(st)-1] == '(') ||
			(v == ']' && st[len(st)-1] == '[') ||
			(v == '}' && st[len(st)-1] == '{')) {
			st = st[:len(st)-1]
		} else {
			return false
		}
	}
	return len(st) == 0
}

func main() {
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("]"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
}
