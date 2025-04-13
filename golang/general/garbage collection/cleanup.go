package garbage

// Using an Explicit Cleanup Method
import (
	"fmt"
	"os"
)

type FileWrapper1 struct {
	file *os.File
}

func (fw *FileWrapper1) Close() {
	if fw.file != nil {
		fmt.Println("Closing file manually")
		fw.file.Close()
	}
}

func cleanupExample() {
	fw := &FileWrapper1{}
	fw.file, _ = os.Open("example.txt")

	// Explicitly calling Close()
	defer fw.Close()
}
