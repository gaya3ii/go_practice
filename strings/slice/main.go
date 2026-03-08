package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func inputText() string {
	fmt.Print("Enter comma-separated values: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// // use package errors to return error
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
