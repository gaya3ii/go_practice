package main

import (
	"bufio"
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
func main() {
	s := inputText()
	if !strings.Contains(s, ",") {
		fmt.Fprintln(os.Stderr, "error: input must contain ',' ")
		os.Exit(1)
	}
	fmt.Println(strings.Split(s, ","))
}
