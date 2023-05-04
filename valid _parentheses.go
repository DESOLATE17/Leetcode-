package main

// https://leetcode.com/problems/valid-parentheses/
// '(', ')', '{', '}', '[' and ']'
// ()[]{}
func isValid(s string) bool {
	check := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]rune, 0)
	size := 0
	for _, v := range []rune(s) {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
			size++
		} else {
			if size == 0 {
				return false
			}

			if val, ok := check[stack[size-1]]; !ok || val != v {
				return false
			}

			stack = stack[:size-1]
			size--
		}
	}
	if size != 0 {
		return false
	}
	return true
}
