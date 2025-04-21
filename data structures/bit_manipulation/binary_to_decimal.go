package bitmanipulation

import "fmt"

// Convert binary string to decimal number
func binaryToDecimal(binaryStr string) (int, error) {
	result := 0
	for _, bit := range binaryStr {
		if bit != '0' && bit != '1' {
			return 0, fmt.Errorf("invalid binary digit: %c", bit)
		}
		result = result * 2
		if bit == '1' {
			result += 1
		}
	}
	return result, nil
}
