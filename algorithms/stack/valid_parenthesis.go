// lc 20: Valid Parentheses
package stack

// IsValid checks for the order of opening and closing brackets
func IsValid(s string) bool {
	pair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != pair[r] {
				return false
			}
		}
	}
	return len(stack) == 0
}
