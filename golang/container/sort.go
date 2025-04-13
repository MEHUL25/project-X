package container

import (
	"fmt"
	"sort"
)

// Function	Description
// sort.Ints([]int)	Sorts a slice of integers in ascending order
// sort.Float64s([]float64)	Sorts a slice of float64s
// sort.Strings([]string)	Sorts a slice of strings

// sort.Sort(sort.Reverse(sort.IntSlice(nums))) - Reversed order

// Functions to Check If Sorted
// Function	Description
// sort.IntsAreSorted([]int)	Checks if ints are sorted
// sort.Float64sAreSorted([]float64)	Checks float sorting
// sort.StringsAreSorted([]string)	Checks string sorting

// Sorting Custom Structs
type Person struct {
	Name string
	Age  int
}

// Sort by Age (ascending)
// Method 1: Define sort.Interface
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func sortExample1() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	sort.Sort(ByAge(people))
	fmt.Println(people)
}

// /Method 2
func sortExample2() {

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// Method 2: sort.Slice (shortcut, modern way)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	// Descending sort:
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})

	// Sort by multiple fields (age, then name):
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Name < people[j].Name
		}
		return people[i].Age < people[j].Age
	})

	// Custom Type Reverse Sort
	sort.Sort(sort.Reverse(ByAge(people)))

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})

}

//  if you're using sort.Sort() with a custom type, then you must define:
//  Len() int
// Swap(i, j int)
// Less(i, j int) bool

// Because sort.Sort() expects a type that implements the sort.Interface:
// type Interface interface {
//     Len() int
//     Less(i, j int) bool
//     Swap(i, j int)
// }
// If you use sort.Slice(), then you do not need to define those methods manually.
// Behind the scenes, Go creates a temporary type that implements Len, Swap, and Less for you using reflection and function closures.
