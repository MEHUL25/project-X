package general

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

// 1. Generic Function: Works with any type that supports addition.
func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

// 2. Generic Struct
type Box[T any] struct {
	value T
}

// Method for Box with a generic type
func (b *Box[T]) GetValue() T {
	return b.value
}

// 3. Interface with Generics: Defines a constraint for types that can be multiplied.
type Multipliable[T any] interface {
	Multiply(T) T
}

// 4. Generic Struct with an Interface Constraint
type Number[T constraints.Ordered] struct {
	val T
}

// // Method for Number struct that multiplies its value
// func (n Number[T]) Multiply(factor T) T {
// 	return n.val * factor
// }

// 5. Generic Function with Multiple Type Parameters
func Compare[T constraints.Ordered, U constraints.Ordered](a T, b U) string {
	if a == any(b).(T) {
		return "Equal"
	}
	return "Not Equal"
}

// 6. Type Embedding with Generics
type Container[T any] struct {
	Item T
}

type SpecialBox[T constraints.Ordered] struct {
	Container[T] // Embedding a generic struct
}

func genericsMain() {
	// Using Generic Function
	fmt.Println("Add Integers:", Add(3, 5))
	fmt.Println("Add Floats:", Add(3.5, 2.5))

	// Using Generic Struct
	intBox := Box[int]{value: 100}
	fmt.Println("Box Value:", intBox.GetValue())

	// Using Generic Interface Implementation
	// num := Number[int]{val: 10}
	// fmt.Println("Multiply:", num.Multiply(2))

	// Using Multiple Type Parameters
	fmt.Println("Compare Int and Float:", Compare(5, 5.0))
	fmt.Println("Compare Strings:", Compare("hello", "world"))

	// Using Type Embedding with Generics
	sBox := SpecialBox[int]{Container[int]{Item: 42}}
	fmt.Println("SpecialBox Value:", sBox.Item)
}
