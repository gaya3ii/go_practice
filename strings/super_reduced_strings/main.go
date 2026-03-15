// In each operation, select a pair of adjacent letters that match, and delete them.
// Delete as many characters as possible using this method and return the length of resulting string.
package main

import (
	"fmt"
)

func superReducedStrings(s string) int {
	stack := []rune{}
	for _, r := range s {
		if len(stack) > 0 && stack[len(stack)-1] == r {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack)

}
func main() {
	s := superReducedStrings("abbcc")
	fmt.Println(s)
}
