package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		key := sortStr(str)
		m[key] = append(m[key], str)

	}

	result := [][]string{}
	for _, group := range m {
		result = append(result, group)
	}

	return result

}

func main() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(result)
}

func sortStr(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
