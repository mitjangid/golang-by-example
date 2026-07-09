# File Handling Examples

Run this file from the project root:

```bash
go run files/files.go
```

You can also run it from inside the folder:

```bash
cd files
go run files.go
```

## What To Practice

- Open, read, write, rename, and delete files.
- Read file metadata.
- Create folders safely.
- Join paths in a cross-platform way.
- Close files after opening them.

## Useful Predefined Functions

```go
os.Open(path) // Open a file for reading.
os.ReadFile(path) // Read a whole small or medium file into memory.
os.WriteFile(path, data, 0644) // Create or replace a file.
os.OpenFile(path, flags, 0644) // Open a file with custom flags such as append.
os.Stat(path) // Read file or folder metadata.
os.MkdirAll(path, 0755) // Create a folder and missing parent folders.
os.ReadDir(path) // List folder entries.
os.Rename(oldPath, newPath) // Move or rename a file.
os.Remove(path) // Delete one file.
filepath.Join("files", "example.json") // Build a path for the current OS.
io.EOF // Check whether reading reached the end of a file.
```

Use this folder when you want to practice real file and folder operations.
