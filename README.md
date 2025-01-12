# Word Count Tool
A simple command-line tool that counts lines, words, bytes, and characters in text files, similar to the Unix wc command.

## Usage
bashCopy./wordcount [options] [filename]

## Options

-l Count lines
-w Count words
-b Count bytes
-m Count characters

If no options are specified, the program defaults to counting lines, words, and bytes.

## Examples
Count everything in a file (defaults to lines, words, and bytes):
```bash
./gowc myfile.txt
7      23     149 myfile.txt
```

### Count only lines:

```bash
./wordcount -l myfile.txt
7 myfile.txt
```
### Count both lines and characters:
```bash
./wordcount -l -m myfile.txt
7      152 myfile.txt
```

## Building from Source
With Go installed, clone this repository and run:
```bash
go build
```
This will create an executable file that you can run directly from the command line.
