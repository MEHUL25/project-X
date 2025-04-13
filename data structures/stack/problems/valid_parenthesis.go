package problems

func isValidParenthesis(s string) bool {
	stack := []rune{}                                         // Stack to store opening brackets
	bracketMap := map[rune]rune{')': '(', '}': '{', ']': '['} // Mapping of closing to opening brackets

	for _, char := range s {
		// If it's a closing bracket
		if open, found := bracketMap[char]; found {
			// Check if stack is empty or top of stack does not match the opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// Pop from stack
			stack = stack[:len(stack)-1]
		} else { // If it's an opening bracket, push onto stack
			stack = append(stack, char)
		}
	}

	// If stack is empty, it's a valid string
	return len(stack) == 0
}
