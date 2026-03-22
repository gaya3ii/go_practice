package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// checks for the order of opening and closing brackets
func isValid(s string) bool {
	pair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != pair[r] {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Print("Enter string with () [] {}: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if input == "" {
		fmt.Fprintln(os.Stderr, "error: empty input")
		return
	}
	fmt.Println(isValid(input))
}
