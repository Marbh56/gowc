package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	var shouldCountBytes bool

	flag.BoolVar(&shouldCountBytes, "b", false, "Count the Bytes of the input")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Error: Please provide a filename\n")
		flag.Usage()
		os.Exit(1)
	}

	filename := args[0]

	if shouldCountBytes {
		count, err := countBytes(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error counting bytes: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d\n", count)
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
