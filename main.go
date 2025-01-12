package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var shouldCountBytes, shouldCountLines, shouldCountWords bool

	flag.BoolVar(&shouldCountBytes, "b", false, "Count bytes")
	flag.BoolVar(&shouldCountLines, "l", false, "Count lines")
	flag.BoolVar(&shouldCountWords, "w", false, "Count words")
	flag.Parse()

	if !shouldCountBytes && !shouldCountLines && !shouldCountWords {
		shouldCountBytes = true
		shouldCountLines = true
		shouldCountWords = true
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Error: Please provide a filename\n")
		flag.Usage()
		os.Exit(1)
	}

	counts, err := processFile(args[0], shouldCountBytes, shouldCountLines, shouldCountWords)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing file: %v\n", err)
		os.Exit(1)
	}

	if shouldCountBytes {
		fmt.Printf("%d bytes\n", counts.bytes)
	}
	if shouldCountLines {
		fmt.Printf("%d lines\n", counts.lines)
	}
	if shouldCountWords {
		fmt.Printf("%d words\n", counts.words)
	}
}

type Counts struct {
	bytes int
	lines int
	words int
}

func processFile(filename string, countBytes, countLines, countWords bool) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	counts := Counts{}

	buffer := make([]byte, 1024)

	inWord := false

	for {
		bytesRead, err := file.Read(buffer)

		if countBytes {
			counts.bytes += bytesRead
		}

		for i := 0; i < bytesRead; i++ {

			if countLines && buffer[i] == '\n' {
				counts.lines++
			}

			if countWords {
				isWhitespace := buffer[i] == ' ' || buffer[i] == '\n' ||
					buffer[i] == '\t' || buffer[i] == '\r'

				if isWhitespace {
					inWord = false
				} else if !inWord {
					counts.words++
					inWord = true
				}
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return Counts{}, err
		}
	}

	return counts, nil
}
