package hashing

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, char := range s {
		m[string(char)]++
	}
	return m
}
