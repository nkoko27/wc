package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"strings"
	"io"
	"bufio"
	"bytes"
)

func main() {

	wcByte := flag.String("c", "stdin", "count bytes in file")
	wcLine := flag.String("l", "", "count lines in a file")
	wcWord := flag.String("w", "", "count words in a file")
	wcChar := flag.String("m", "", "count characters in a file")

	switch os.Args[1] {
	case "-c":
		if len(os.Args) > 2 {
			flag.Parse()
			count := byteCount(*wcByte)
			fmt.Printf("%d %q\n", count, *wcByte)
		} else {
			stdinByteCount()
		}
	case "-l":
		if len(os.Args) > 2 {
			flag.Parse()
			count := lineCount(*wcLine)
			fmt.Printf("%d %q\n", count, *wcLine)
		} else {
			stdinLineCount()
		}
	case "-w":
		if len(os.Args) > 2 {
			flag.Parse()
			count := wordCount(*wcWord)
			fmt.Printf("%d %q\n", count, *wcWord)
		} else {
			stdinWordCount()
		}
		
	case "-m":
		if len(os.Args) > 2{
			flag.Parse()
			count := charCount(*wcChar)
			fmt.Printf("%d %q\n", count, *wcChar)
		} else {
			stdinCharCount()
		}
	default:
		flag.Parse()
		bc := byteCount(os.Args[1])
		lc := lineCount(os.Args[1])
		wc := wordCount(os.Args[1])
		fmt.Printf("%d %d %d %q\n", bc, lc, wc, os.Args[1])
	}

}

func charCount(text string) int {
	file, err := os.Open(text)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	chars := []rune(string(data))
	return len(chars)
}

func stdinCharCount() {
	reader := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	// data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n")) // Only necessary if running from powershell
	chars := []rune(string(data))
	fmt.Printf("%d", len(chars))
}

func wordCount(text string) int {
	file, err := os.Open(text)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, int(stats.Size()))
	_, err = file.Read(data) // look at io.ReadAll

	s := string(data[:])

	words := strings.Fields(s)	
	return len(words)
}

func stdinWordCount() {
	reader := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	s := string(data[:])
	words := strings.Fields(s)
	fmt.Printf("%d", len(words))
}

func lineCount(text string) int {
	file, err := os.Open(text)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, int(stats.Size()))
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	line := 0
	for _, val := range data {
		if string(val) == "\n" {
			line++
		}
	}
	return line
}

func stdinLineCount() {
    reader := bufio.NewReader(os.Stdin)
    totalLines := 0
    for {
        _, err := reader.ReadBytes('\n')
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        totalLines++
    }
    fmt.Printf("Total lines: %d\n", totalLines)
}

func byteCount(text string) int {
	file, err := os.Open(text)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	length := stats.Size()
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, length)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func stdinByteCount() {
	reader := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	// data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n")) // only necessary if running from powershell
	fmt.Printf("%d", len(data))
}