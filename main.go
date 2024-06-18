package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func byteCounter(filePath string) {
	fi, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fi.Size())
}
func lineCounter(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// initialise an empty buffer to hold the file contents
	buf := make([]byte, 32*1024)
	// variable to hold line count
	count := 0
	// slice containing single byte representing newline character
	lineSep := []byte{'\n'}

	// infinite loop
	for {
		// read *buffer_size* bytes from the file into the buffer, returning the number of bytes read(c)
		c, err := file.Read(buf)
		// count number of newline characters in the bytes that have been read
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			fmt.Println(count)
			return

		case err != nil:
			log.Fatal(err)
			return
		}
	}
}

func countWords(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := make([]byte, 32*1024)
	count := 0
	inWord := false
	for {
		c, err := file.Read(buf)
		for i := 0; i < c; i++ {
			if buf[i] == ' ' || buf[i] == '\n' || buf[i] == '\t' || buf[i] == '\r' {
				if inWord {
					count++
				}
				inWord = false
			} else {
				inWord = true
			}
		}
		if err == io.EOF {
			if inWord {
				count++
			}
			fmt.Println(count)
			return
		}
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
func countBytes(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fileContents := string(data)
	fmt.Println(len(fileContents))
}

func main() {
	cPtr := flag.Bool("c", false, "display number of bytes of file")
	lPtr := flag.Bool("l", false, "display number of lines")
	wPtr := flag.Bool("w", false, "display number of words")
	mPtr := flag.Bool("m", false, "display number of characters")
	filePath := os.Args[len(os.Args)-1]

	masterFlag := !*cPtr && !*lPtr && !*wPtr && !*mPtr

	if masterFlag {
		*cPtr = true
		*lPtr = true
		*wPtr = true
		*mPtr = true
	}

	flag.Parse()

	if *cPtr {
		byteCounter(filePath)
	}
	if *lPtr {
		lineCounter(filePath)
	}
	if *wPtr {
		countWords(filePath)
	}
	if *mPtr {
		countBytes(filePath)
	}
}
