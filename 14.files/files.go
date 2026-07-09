package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()

	// These helpers make the example work from both:
	// go run files/files.go
	// cd files && go run files.go
	exampleFile := findPath("files/example.json", "example.json")
	exampleFolder := findPath("files/exp_folder", "exp_folder")

	// Open returns an *os.File for reading. Always close files you open.
	file, err := os.Open(exampleFile)
	if err != nil {
		fmt.Println("open file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File name:", file.Name())

	// Stat returns metadata such as size, permissions, and modified time.
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("file stat:", err)
		return
	}
	fmt.Println("File size:", fileInfo.Size(), "bytes")
	fmt.Println("File permissions:", fileInfo.Mode())

	// Read reads from the current file position.
	buffer := make([]byte, 80)
	bytesRead, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("read file:", err)
		return
	}
	fmt.Printf("First %d bytes:\n%s\n", bytesRead, buffer[:bytesRead])

	// Seek changes the current position. Here we go back to the start.
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		fmt.Println("seek file:", err)
		return
	}

	// ReadAt reads from a specific offset without changing the current position.
	readAtBuffer := make([]byte, 40)
	bytesReadAt, err := file.ReadAt(readAtBuffer, 10)
	if err != nil && err != io.EOF {
		fmt.Println("read at offset:", err)
		return
	}
	fmt.Printf("Bytes from offset 10:\n%s\n", readAtBuffer[:bytesReadAt])

	// os.ReadFile is the simple modern way to read a whole file into memory.
	// Use it for small/medium files. For very large files, stream with os.Open.
	content, err := os.ReadFile(exampleFile)
	if err != nil {
		fmt.Println("read whole file:", err)
		return
	}
	fmt.Println("Whole file loaded bytes:", len(content))

	// os.Stat also works directly with a path.
	pathInfo, err := os.Stat(exampleFile)
	if err != nil {
		fmt.Println("path stat:", err)
		return
	}
	fmt.Println("Is directory:", pathInfo.IsDir())

	// os.MkdirAll creates a folder and any missing parent folders.
	if err := os.MkdirAll(exampleFolder, 0755); err != nil {
		fmt.Println("create folder:", err)
		return
	}

	// os.WriteFile creates or replaces a file with the provided bytes.
	// The process id keeps the demo filename unique if examples run together.
	demoFile := filepath.Join(exampleFolder, fmt.Sprintf("demo_file_%d.txt", os.Getpid()))
	if err := os.WriteFile(demoFile, []byte("Learning Go file handling\n"), 0644); err != nil {
		fmt.Println("write file:", err)
		return
	}

	// os.OpenFile gives more control. Here we append to the existing file.
	appendFile, err := os.OpenFile(demoFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file for append:", err)
		return
	}
	if _, err := appendFile.WriteString("Appending another line\n"); err != nil {
		appendFile.Close()
		fmt.Println("append file:", err)
		return
	}
	if err := appendFile.Close(); err != nil {
		fmt.Println("close appended file:", err)
		return
	}

	// os.ReadDir lists folder entries.
	entries, err := os.ReadDir(exampleFolder)
	if err != nil {
		fmt.Println("read folder:", err)
		return
	}
	for _, entry := range entries {
		fmt.Println("Folder entry:", entry.Name(), "is directory:", entry.IsDir())
	}

	// os.Rename moves or renames a file.
	renamedFile := filepath.Join(exampleFolder, fmt.Sprintf("renamed_demo_file_%d.txt", os.Getpid()))
	_ = os.Remove(renamedFile) // Ignore the error when the file does not exist.
	if err := os.Rename(demoFile, renamedFile); err != nil {
		fmt.Println("rename file:", err)
		return
	}

	// os.Remove deletes a file. Use os.RemoveAll only when deleting folders.
	if err := os.Remove(renamedFile); err != nil {
		fmt.Println("remove file:", err)
		return
	}

	duration := time.Since(start)
	fmt.Printf("\nCode execution took: %s\n", duration)
}

func findPath(paths ...string) string {
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return paths[0]
}
