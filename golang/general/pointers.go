package general

import (
	"fmt"
	"sync"
)

// Struct with pointer
type Person2 struct {
	Name string
	Age  int
}

// Function modifying value using pointer
func modifyValue(p *int) {
	*p = *p + 10
}

// Function modifying struct using pointer
func modifyPerson(p *Person2) {
	p.Age += 5
}

// Pointer to Pointer example
func doublePointer(pp **int) {
	**pp += 20
}

// Sync.Pool usage to optimize memory reuse
var pool = sync.Pool{
	New: func() interface{} {
		return new(int)
	},
}

func pointerExample() {
	// 1️⃣ Basic Pointer Usage
	num := 42
	ptr := &num
	fmt.Println("Original num:", num)
	fmt.Println("Pointer address:", ptr)
	fmt.Println("Value at pointer:", *ptr)

	// 2️⃣ Passing pointer to function
	modifyValue(ptr)
	fmt.Println("After modifyValue:", num) // Should be 52

	// 3️⃣ Pointer to Struct
	p1 := Person2{"John", 25}
	modifyPerson(&p1)
	fmt.Println("Updated Age:", p1.Age) // 30

	// 4️⃣ Pointer to Array
	arr := [3]int{10, 20, 30}
	pArr := &arr
	fmt.Println("Pointer to array, first element:", (*pArr)[0])

	// 5️⃣ Double Pointer Example
	p2 := &num
	pp := &p2
	doublePointer(pp)
	fmt.Println("After doublePointer, num:", num) // 72

	// 6️⃣ Using new() to allocate memory
	pNew := new(int) // Zero-initialized memory
	*pNew = 100
	fmt.Println("Value at new pointer:", *pNew) // 100

	// 7️⃣ Checking for nil pointers
	var pNil *int
	if pNil == nil {
		fmt.Println("pNil is nil")
	}

	// 8️⃣ Using sync.Pool
	pooledNum := pool.Get().(*int)
	*pooledNum = 10
	fmt.Println("Pooled Number:", *pooledNum)
	pool.Put(pooledNum) // Return to pool for reuse
}
