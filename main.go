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

	if len(args) == 0 {
		counts, err := processReader(os.Stdin, shouldCountBytes, shouldCountLines, shouldCountWords, shouldCountCharacters)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing stdin: %v\n", err)
			os.Exit(1)
		}
		printCounts(counts, "")
		return
	}
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening %s: %v\n", filename, err)
			continue
		}

		counts, err := processReader(file, shouldCountBytes, shouldCountLines, shouldCountWords, shouldCountCharacters)
		file.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", filename, err)
			continue
		}

		printCounts(counts, filename)
	}
}

type Counts struct {
	bytes      int
	lines      int
	words      int
	characters int
}

func processReader(reader io.Reader, countBytes, countLines, countWords, countCharacters bool) (Counts, error) {
	bufReader := bufio.NewReader(reader)
	counts := Counts{}
	inWord := false

	for {
		r, size, err := bufReader.ReadRune()
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

func printCounts(counts Counts, filename string) {
	var output string

	if counts.lines > 0 {
		output += fmt.Sprintf(" %7d", counts.lines)
	}
	if counts.words > 0 {
		output += fmt.Sprintf(" %7d", counts.words)
	}
	if counts.bytes > 0 {
		output += fmt.Sprintf(" %7d", counts.bytes)
	}
	if counts.characters > 0 {
		output += fmt.Sprintf(" %7d", counts.characters)
	}

	if filename != "" {
		output += fmt.Sprintf(" %s", filename)
	}

	fmt.Println(output)
}
