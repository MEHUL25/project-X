package math

// gcd function calculates the greatest common divisor of two numbers using the Euclidean algorithm
func gcd(a, b int) int {
	// While b is not zero, continue applying the Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
