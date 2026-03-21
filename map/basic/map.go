package main

import "fmt"

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, char := range s {
		m[string(char)]++
	}
	return m
}

func main() {
	s := "golang"
	fmt.Println(wordCount(s))
}
