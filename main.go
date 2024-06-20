package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func byteCounterForFilePath(filePath string) {
	fi, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fi.Size())
}
func byteCounterForReader(reader io.Reader) {
	buf := make([]byte, 32*1024)
	totalBytes := 0

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		totalBytes += n
	}

	fmt.Println(totalBytes)
}

func lineCounter(file io.Reader) {
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
			// reset file pointer to beginning of file
			//_, err := file.(io.Seeker).Seek(0, 0)
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			return

		case err != nil:
			log.Fatal(err)
			return
		}
	}
}

func countWords(file io.Reader) {
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

func countCharactersForFilePath(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fileContents := string(data)
	fmt.Println(len(fileContents))
}
func countCharactersForReader(reader io.Reader) {
	buf := make([]byte, 32*1024)
	totalChars := 0

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		totalChars += len(string(buf[:n]))
	}

	fmt.Println(totalChars)
}

func main() {
	cPtr := flag.Bool("c", false, "display number of bytes of file")
	lPtr := flag.Bool("l", false, "display number of lines")
	wPtr := flag.Bool("w", false, "display number of words")
	mPtr := flag.Bool("m", false, "display number of characters")
	filePath := flag.String("path", "", "file path")
	flag.Parse()

	masterFlag := !*cPtr && !*lPtr && !*wPtr && !*mPtr

	if masterFlag {
		*cPtr = true
		*lPtr = true
		*wPtr = true
	}

	var fileReader io.Reader
	var err error
	var fileBytes []byte

	if *filePath == "" {
		fileBytes, err = io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fileReader = bytes.NewReader(fileBytes)
	} else {
		fileReader, err = os.Open(*filePath)
		if err != nil {
			log.Fatal(err)
		}
	}
	if closer, ok := fileReader.(io.Closer); ok {
		defer closer.Close()
	}

	if *cPtr {
		if *filePath != "" {
			byteCounterForFilePath(*filePath)
		} else {
			byteCounterForReader(fileReader)
			fileReader = bytes.NewReader(fileBytes)
		}
	}
	if *lPtr {
		lineCounter(fileReader)
		if *filePath == "" {
			fileReader = bytes.NewReader(fileBytes)
		} else {
			resetFilePointer(fileReader)
		}
	}
	if *wPtr {
		countWords(fileReader)
		if *filePath == "" {
			fileReader = bytes.NewReader(fileBytes)
		} else {
			resetFilePointer(fileReader)
		}
	}
	if *mPtr {
		if *filePath != "" {
			countCharactersForFilePath(*filePath)
		} else {
			countCharactersForReader(fileReader)
			fileReader = bytes.NewReader(fileBytes)
		}
	}
}

func resetFilePointer(file io.Reader) {
	_, err := file.(io.Seeker).Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
}
