// Read comma-separated input from stdin, validate it contains a comma,
// then split and print the resulting slice.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// inputText prompts the user and reads one trimmed line from stdin.
func inputText() string {
	fmt.Print("Enter comma-separated values: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// splitInput splits on comma and returns the slice, or an error if there is no comma.
func splitInput(s string) ([]string, error) {
	if !strings.Contains(s, ",") {
		return nil, errors.New("error: input must contain ',' as separator")
	}
	return strings.Split(s, ","), nil
}

func main() {
	s := inputText()
	result, err := splitInput(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(result)
}
