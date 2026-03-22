// camelcase counts how many words are in a camelCase string.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func camelcase(s string) (int32, error) {
	if s == "" {
		return 0, errors.New("empty input")
	}
	count := 1
	for _, r := range s {
		if unicode.IsUpper(r) {
			count++
		}
	}
	return int32(count), nil
}


func main() {
	fmt.Print("Enter a camelCase string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	count, err := camelcase(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Println(count)
}
