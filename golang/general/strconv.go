package general

import (
	"fmt"
	"strconv"
)

// Prominent strconv Functions (Grouped by Type)
// 🔢 Integer Conversions
// Function	Description
// strconv.Itoa(int)	Converts int to string
// strconv.Atoi(string)	Converts string to int
// strconv.FormatInt(int64, base)	Converts int64 to string in given base (e.g., 2, 10, 16)
// strconv.ParseInt(string, base, bitSize)	Parses string into int64, int32, etc.
// strconv.FormatUint(uint64, base)	Converts uint64 to string
// strconv.ParseUint(string, base, bitSize)	Parses string to uint64, etc.
// 💡 Boolean Conversions
// Function	Description
// strconv.FormatBool(bool)	Converts bool to "true" or "false"
// strconv.ParseBool(string)	Parses "1", "t", "T", "true", "TRUE" → true (case insensitive)
// 🧮 Float Conversions
// Function	Description
// strconv.FormatFloat(float64, fmt, prec, bitSize)	Converts float64 to string with control over formatting
// strconv.ParseFloat(string, bitSize)	Converts string to float (float32 or float64)
// 🧪 Quote and Unquote Strings
// Function	Description
// strconv.Quote(string)	Adds double quotes + escapes special characters
// strconv.QuoteToASCII(string)	Like Quote, but escapes non-ASCII characters
// strconv.Unquote(string)	Removes quotes and unescapes ("abc\n" → abc<newline>)
// strconv.QuoteRune(rune)	Converts a rune to quoted string form
// strconv.UnquoteChar()	Used to read a single quoted character (low-level, less common)
// 🔍 Checking and Conversions Helpers
// Function	Description
// strconv.IsPrint(rune)	Returns true if rune is printable

func main() {
	fmt.Println("🔢 Integer Conversions:")
	// int to string
	str := strconv.Itoa(42)
	fmt.Println("Itoa(42):", str)

	// string to int
	i, _ := strconv.Atoi("42")
	fmt.Println("Atoi(\"42\"):", i)

	// int64 to binary/decimal/hex string
	fmt.Println("FormatInt(255, 2):", strconv.FormatInt(255, 2))   // binary
	fmt.Println("FormatInt(255, 10):", strconv.FormatInt(255, 10)) // decimal
	fmt.Println("FormatInt(255, 16):", strconv.FormatInt(255, 16)) // hex

	// string to int64
	val64, _ := strconv.ParseInt("1010", 2, 64)
	fmt.Println("ParseInt(\"1010\", 2, 64):", val64)

	// uint64 to string
	fmt.Println("FormatUint(123, 10):", strconv.FormatUint(123, 10))

	// string to uint64
	uval64, _ := strconv.ParseUint("123", 10, 64)
	fmt.Println("ParseUint(\"123\", 10, 64):", uval64)

	// -------------------------------------

	fmt.Println("\n💡 Boolean Conversions:")
	fmt.Println("FormatBool(true):", strconv.FormatBool(true))

	b, _ := strconv.ParseBool("t")
	fmt.Println("ParseBool(\"t\"):", b)

	// -------------------------------------

	fmt.Println("\n🧮 Float Conversions:")
	// float64 to string
	fmt.Println("FormatFloat(3.1415, 'f', 2, 64):", strconv.FormatFloat(3.1415, 'f', 2, 64)) // "3.14"

	// string to float64
	fval, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println("ParseFloat(\"3.1415\", 64):", fval)

	// -------------------------------------

	fmt.Println("\n🧪 Quote and Unquote:")
	quoted := strconv.Quote("Hello\nGopher")
	fmt.Println("Quote(\"Hello\\nGopher\"):", quoted)

	asciiQuoted := strconv.QuoteToASCII("Gølang ♥")
	fmt.Println("QuoteToASCII(\"Gølang ♥\"):", asciiQuoted)

	unquoted, _ := strconv.Unquote("\"Hello\\nGopher\"")
	fmt.Println("Unquote(\"\\\"Hello\\\\nGopher\\\"\"):", unquoted)

	runeQuoted := strconv.QuoteRune('❤')
	fmt.Println("QuoteRune('❤'):", runeQuoted)

	// -------------------------------------

	fmt.Println("\n🔍 Miscellaneous:")
	fmt.Println("IsPrint('A'):", strconv.IsPrint('A'))    // true
	fmt.Println("IsPrint('\\n'):", strconv.IsPrint('\n')) // false
}

// Sample Output:

// 🔢 Integer Conversions:
// Itoa(42): 42
// Atoi("42"): 42
// FormatInt(255, 2): 11111111
// FormatInt(255, 10): 255
// FormatInt(255, 16): ff
// ParseInt("1010", 2, 64): 10
// FormatUint(123, 10): 123
// ParseUint("123", 10, 64): 123

// 💡 Boolean Conversions:
// FormatBool(true): true
// ParseBool("t"): true

// 🧮 Float Conversions:
// FormatFloat(3.1415, 'f', 2, 64): 3.14
// ParseFloat("3.1415", 64): 3.1415

// 🧪 Quote and Unquote:
// Quote("Hello\nGopher"): "Hello\nGopher"
// QuoteToASCII("Gølang ♥"): "G\u00f8lang \u2665"
// Unquote("\"Hello\\nGopher\""): Hello
// Gopher
// QuoteRune('❤'): '❤'

// 🔍 Miscellaneous:
// IsPrint('A'): true
// IsPrint('\n'): false
