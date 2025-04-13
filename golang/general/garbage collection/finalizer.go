package garbage

// Using a Finalizer to Close a File

import (
	"fmt"
	"os"
	"runtime"
)

type FileWrapper struct {
	file *os.File
}

func finzalizerExample() {
	fw := &FileWrapper{}
	fw.file, _ = os.Open("example.txt")

	// Attach a finalizer to ensure file is closed before GC
	runtime.SetFinalizer(fw, func(f *FileWrapper) {
		fmt.Println("Finalizer: Closing file")
		f.file.Close()
	})

	// Remove reference to trigger GC
	fw = nil

	// Force GC (only for demo, normally avoid this)
	runtime.GC()
}

// If GC runs and fw is unreachable, the finalizer will close the file.
