package problems

import (
	"fmt"
)

// isBalanced function to check if brackets are balanced
func IsBalanced(s string) string {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1] // Pop from stack
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func RunBalancedBracketTests() {
	tests := []string{
		"{[()]}",
		"{[(])}",
		"{(([])[[]])[]}",
	}

	for _, test := range tests {
		fmt.Printf("Input: %s => Output: %s\n", test, IsBalanced(test))
	}
}
