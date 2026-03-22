package main

import "fmt"

func nonRepeat(s string) string {
	m := make(map[string]int)
	for _, ch := range s {
		m[string(ch)]++
	}
	for _, ch := range s {
		if m[string(ch)] == 1 {
			return string(ch)
		}
	}
	return ""
}
func main() {
	s := "leetcode"
	fmt.Println(nonRepeat(s))

}
