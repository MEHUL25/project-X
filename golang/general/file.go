package general

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileOperationsDemo() {
	filename := "example.txt"

	// Create and write to a file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	file.WriteString("Hello, Go File Operations!\n")
	fmt.Fprintln(file, "Another line in the file.")

	// Read entire file content
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	fmt.Println("File content:")
	fmt.Println(string(data))

	// Read entire file content - method 2
	data, err = os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	fmt.Println("File content:")
	fmt.Println(string(data))

	// Append to file
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file for appending: %v", err)
	}
	defer file.Close()
	file.WriteString("Appended text\n")

	// Read file line by line
	file, err = os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	fmt.Println("Reading file line by line:")
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatalf("Failed to read file: %v", err)
		}
		if n == 0 {
			break
		}
		fmt.Print(string(buf[:n]))
	}

	// Delete the file
	err = os.Remove(filename)
	if err != nil {
		log.Fatalf("Failed to delete file: %v", err)
	}
	fmt.Println("File deleted successfully.")
}
