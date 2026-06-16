// In each operation, select a pair of adjacent letters that match, and delete them.
// Delete as many characters as possible using this method and return the length of resulting string.
package stack

func SuperReducedStrings(s string) int {
	st := []rune{}
	for _, r := range s {
		if len(st) > 0 && st[len(st)-1] == r {
			st = st[:len(st)-1]
		} else {
			st = append(st, r)
		}
	}
	return len(st)
}
