package main

import "fmt"

var opening = map[rune]bool{
	'(': true,
	'[': true,
	'{': true,
}

var closing = map[rune]bool{
	')': true,
	']': true,
	'}': true,
}

var matching = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func BalancedBrackets(s string) bool {

	stack := []rune{}
	for _, char := range s {

		if _, ok := opening[char]; ok {
			stack = append(stack, char)
			continue
		}
		if _, ok := closing[char]; ok {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == matching[char] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	str := "([])(){}(())()()"
	fmt.Println(BalancedBrackets(str))

}
