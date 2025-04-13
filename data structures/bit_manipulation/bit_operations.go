package bitmanipulation

import (
	"fmt"
)



func main() {
	var a uint8 = 12 // 00001100
	var b uint8 = 10 // 00001010

	fmt.Printf("a: %08b (%d)\n", a, a)
	fmt.Printf("b: %08b (%d)\n\n", b, b)

	// Bitwise AND
	and := a & b
	fmt.Printf("a & b  : %08b (%d)\n", and, and)

	// Bitwise OR
	or := a | b
	fmt.Printf("a | b  : %08b (%d)\n", or, or)

	// Bitwise XOR
	xor := a ^ b
	fmt.Printf("a ^ b  : %08b (%d)\n", xor, xor)

	// Bitwise AND NOT (a &^ b)
	andNot := a &^ b
	fmt.Printf("a &^ b : %08b (%d)\n", andNot, andNot)

	// Bitwise NOT (on a)
	notA := ^a
	fmt.Printf("^a     : %08b (%d)\n", notA, notA)

	// Left shift (a << 2)
	leftShift := a << 2
	fmt.Printf("a << 2 : %08b (%d)\n", leftShift, leftShift)

	// Right shift (a >> 2)
	rightShift := a >> 2
	fmt.Printf("a >> 2 : %08b (%d)\n", rightShift, rightShift)
}
