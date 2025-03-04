package problems

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func RunBalancedBracketTests() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan string bracket: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := IsBalanced(input)
	fmt.Println("Output:", result)
}
