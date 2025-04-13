package container

import (
	"fmt"
	"strings"
)

func main() {
	base := "  Hello, GoLang World!  "
	clean := strings.TrimSpace(base) // remove leading/trailing space

	// 🔤 Basic Functions
	fmt.Println("🔤 Basic Functions:")
	fmt.Println("Contains:", strings.Contains(clean, "GoLang"))
	fmt.Println("HasPrefix:", strings.HasPrefix(clean, "Hello"))
	fmt.Println("HasSuffix:", strings.HasSuffix(clean, "World!"))
	fmt.Println("Count of 'l':", strings.Count(clean, "l"))
	fmt.Println("Index of 'GoLang':", strings.Index(clean, "GoLang"))
	fmt.Println("LastIndex of 'l':", strings.LastIndex(clean, "l"))

	// ✂️ Splitting & Joining
	fmt.Println("\n✂️ Splitting & Joining:")
	words := strings.Split(clean, " ")
	fmt.Println("Split:", words)
	fmt.Println("Join with '-':", strings.Join(words, "-"))
	fmt.Println("SplitN (2 parts):", strings.SplitN(clean, " ", 2))

	// 🔤 Case Conversion
	fmt.Println("\n🔤 Case Conversion:")
	fmt.Println("ToLower:", strings.ToLower(clean))
	fmt.Println("ToUpper:", strings.ToUpper(clean))
	// Title is deprecated but still works:
	fmt.Println("Title (deprecated):", strings.Title(clean)) // ⚠️ Deprecated

	// 🔄 Replacing & Trimming
	fmt.Println("\n🔄 Replacing & Trimming:")
	fmt.Println("Replace first 'l' with 'L':", strings.Replace(clean, "l", "L", 1))
	fmt.Println("ReplaceAll 'l' with 'L':", strings.ReplaceAll(clean, "l", "L"))
	fmt.Println("Trim '#' from '#Hello#':", strings.Trim("#Hello#", "#"))
	fmt.Println("TrimPrefix 'Hello' from 'HelloWorld':", strings.TrimPrefix("HelloWorld", "Hello"))
	fmt.Println("TrimSuffix 'World' from 'HelloWorld':", strings.TrimSuffix("HelloWorld", "World"))

	// 🧪 Comparison
	fmt.Println("\n🧪 Comparison:")
	fmt.Println("EqualFold('GoLang', 'golang'):", strings.EqualFold("GoLang", "golang"))

	// 🧵 Builder
	fmt.Println("\n🧵 Builder:")
	var builder strings.Builder
	builder.WriteString("Go")
	builder.WriteString("Lang")
	fmt.Println("Built string:", builder.String())
}

// Output :

// 🔤 Basic Functions:
// Contains: true
// HasPrefix: true
// HasSuffix: true
// Count of 'l': 3
// Index of 'GoLang': 7
// LastIndex of 'l': 20

// ✂️ Splitting & Joining:
// Split: [Hello, GoLang World!]
// Join with '-': Hello,-GoLang-World!
// SplitN (2 parts): [Hello, GoLang World!]

// 🔤 Case Conversion:
// ToLower: hello, golang world!
// ToUpper: HELLO, GOLANG WORLD!
// Title (deprecated): Hello, Golang World!

// 🔄 Replacing & Trimming:
// Replace first 'l' with 'L': HeLlo, GoLang World!
// ReplaceAll 'l' with 'L': HeLLo, GoLang WorLd!
// Trim '#' from '#Hello#': Hello
// TrimPrefix 'Hello' from 'HelloWorld': World
// TrimSuffix 'World' from 'HelloWorld': Hello

// 🧪 Comparison:
// EqualFold('GoLang', 'golang'): true

// 🧵 Builder:
// Built string: GoLang
