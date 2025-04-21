package bitmanipulation

import "strconv"

// Convert decimal number to binary string
func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	binary := ""
	for n > 0 {
		remainder := n % 2
		binary = strconv.Itoa(remainder) + binary
		n = n / 2
	}
	return binary
}
