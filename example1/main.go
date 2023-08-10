package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	words := CountWords(os.Args[1])
	fmt.Printf("%q: %d words\n", os.Args[1], words)
}

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func CountWords(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}

	words := 0
	inword := false
	for {
		r, err := readbyte(f)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}

	return words
}
