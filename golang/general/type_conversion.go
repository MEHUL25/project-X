package general

import (
	"fmt"
	"strconv"
)

func typeConversionExamples() {
	// Example 1: Integer to Float
	intValue := 42
	floatValue := float64(intValue)
	fmt.Println("Integer to Float:", floatValue)

	// Example 2: Float to Integer
	floatNum := 42.9
	intNum := int(floatNum) // Truncates the decimal part
	fmt.Println("Float to Integer:", intNum)

	// Example 3: Integer to String
	strValue := strconv.Itoa(intValue)
	fmt.Println("Integer to String:", strValue)

	// Example 4: String to Integer
	strNum := "123"
	parsedInt, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Println("String to Integer:", parsedInt)
	}

	// Example 5: String to Float
	strFloat := "123.45"
	parsedFloat, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Println("String to Float:", parsedFloat)
	}

	// Example 5.1: Float to String
	f := 3.14159
	s := strconv.FormatFloat(f, 'f', 2, 64) // 'f' format, 2 decimal places, 64-bit float
	fmt.Println(s)                          // Output: "3.14"

	// Example 6: Boolean to String
	boolValue := true
	strBool := strconv.FormatBool(boolValue)
	fmt.Println("Boolean to String:", strBool)

	// Example 7: String to Boolean
	boolStr := "true"
	parsedBool, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("Error converting string to bool:", err)
	} else {
		fmt.Println("String to Boolean:", parsedBool)
	}

	// Example 8: Byte Slice to String
	byteSlice := []byte{'H', 'e', 'l', 'l', 'o'}
	stringFromBytes := string(byteSlice)
	fmt.Println("Byte Slice to String:", stringFromBytes)

	// Example 9: String to Byte Slice
	originalString := "Hello"
	byteArray := []byte(originalString)
	fmt.Println("String to Byte Slice:", byteArray)
}
