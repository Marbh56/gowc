package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	var shouldCountBytes bool
	var shouldCountLines bool
	var shouldCountWords bool

	flag.BoolVar(&shouldCountBytes, "b", false, "Count the Bytes of the input")
	flag.BoolVar(&shouldCountLines, "l", false, "Count the number of lines in the input")
	flag.BoolVar(&shouldCountWords, "w", false, "Count the number of words in the input")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Error: Please provide a filename\n")
		flag.Usage()
		os.Exit(1)
	}

	filename := args[0]

	if !shouldCountBytes && !shouldCountLines && !shouldCountWords {
		shouldCountBytes = true
		shouldCountLines = true
		shouldCountWords = true
	}
	if shouldCountWords {
		count, err := countWords(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error counting words: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d -- Words\n", count)
	}

	if shouldCountLines {
		count, err := countLines(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error counting lines: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d -- Lines\n", count)
	}

	if shouldCountBytes {
		count, err := countBytes(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error counting bytes: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d -- Bytes\n", count)
	}
}

func countBytes(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	totalBytes := 0

	for {
		bytesRead, err := file.Read(buffer)
		totalBytes += bytesRead

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}
	}

	return totalBytes, nil
}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	lineCount := 0

	for {
		bytesRead, err := file.Read(buffer)

		for i := 0; i < bytesRead; i++ {
			if buffer[i] == '\n' {
				lineCount++
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}
	}
	return lineCount, nil
}

func countWords(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	wordCount := 0
	inWord := false

	for {
		bytesRead, err := file.Read(buffer)

		for i := 0; i < bytesRead; i++ {
			isWhitespace := buffer[i] == ' ' || buffer[i] == '\n' ||
				buffer[i] == '\t' || buffer[i] == '\r'
			if isWhitespace {
				inWord = false
			} else if !inWord {
				wordCount++
				inWord = true
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
	}

	return wordCount, nil

}
