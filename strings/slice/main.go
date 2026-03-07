package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter comma-separated values: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if !strings.Contains(input, ",") {
		fmt.Fprintln(os.Stderr, "error: input must contain ',' ")
		os.Exit(1)
	}
	fmt.Println(strings.Split(input, ","))
}
