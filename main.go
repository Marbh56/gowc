package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var shouldCountBytes, shouldCountLines, shouldCountWords, shouldCountCharacters bool

	flag.BoolVar(&shouldCountBytes, "b", false, "Count bytes")
	flag.BoolVar(&shouldCountLines, "l", false, "Count lines")
	flag.BoolVar(&shouldCountWords, "w", false, "Count words")
	flag.BoolVar(&shouldCountCharacters, "m", false, "Count characters")
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

	counts, err := processFile(args[0], shouldCountBytes, shouldCountLines, shouldCountWords, shouldCountCharacters)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing file: %v\n", err)
		os.Exit(1)
	}

	if shouldCountBytes {
		fmt.Printf("%d -- bytes\n", counts.bytes)
	}
	if shouldCountLines {
		fmt.Printf("%d -- lines\n", counts.lines)
	}
	if shouldCountWords {
		fmt.Printf("%d -- words\n", counts.words)
	}
	if shouldCountCharacters {
		fmt.Printf("%d -- characters\n", counts.characters)
	}
}

type Counts struct {
	bytes      int
	lines      int
	words      int
	characters int
}

func processFile(filename string, countBytes, countLines, countWords, countCharacters bool) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	counts := Counts{}

	inWord := false

	for {
		r, size, err := reader.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			return Counts{}, err
		}

		if countBytes {
			counts.bytes += size
		}

		if countCharacters {
			counts.characters++
		}

		if countLines && r == '\n' {
			counts.lines++
		}

		if countWords {
			isWhitespace := r == ' ' || r == '\n' || r == '\t' || r == '\r'

			if isWhitespace {
				inWord = false
			} else if !inWord {
				counts.words++
				inWord = true
			}
		}
	}
	return counts, nil
}
