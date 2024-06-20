# go-wc 

A simple implementation of the `wc` command in Go. A solution to Code Challenge #1 curated by John Crickett.

## Usage
This implementation supports both file path and standard input.
### File path
```bash
go run main.go -path=<file_path>
```
### Standard input
```bash
cat test.txt | go run main.go
```

This implementation supports the following flags:
- `-l`: print the number of lines in the input
- `-w`: print the number of words in the input
- `-c`: print the number of bytes in the input
- `-m`: print the number of characters in the input (may differ due to locale differences)

If no flags are provided, `-l`, `-w` and `-c` are used by default.