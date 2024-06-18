package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func lineCounter(r io.Reader) (int, error) {
	// initialise an empty buffer to hold the file contents
	buf := make([]byte, 32*1024)
	// variable to hold line count
	count := 0
	// slice containing single byte representing newline character
	lineSep := []byte{'\n'}

	// infinite loop
	for {
		// read *buffer_size* bytes from the file into the buffer, returning the number of bytes read(c)
		c, err := r.Read(buf)
		// count number of newline characters in the bytes that have been read
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func countWords(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	inWord := false
	for {
		c, err := r.Read(buf)
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
			return count, nil
		}
		if err != nil {
			return count, err
		}
	}
}

func main() {
	cPtr := flag.Bool("c", true, "display number of bytes of file")
	lPtr := flag.Bool("l", true, "display number of lines")
	wPtr := flag.Bool("w", true, "display number of words")
	filePath := os.Args[len(os.Args)-1]

	flag.Parse()

	if *cPtr {
		fi, err := os.Stat(filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fi.Size())
	}
	if *lPtr {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		lines, err := lineCounter(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(lines)
	}
	if *wPtr {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		words, err := countWords(file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(words)
	}
}
