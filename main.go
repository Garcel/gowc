package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Mode int

const (
	Words Mode = iota
	Lines
	Bytes
)

func main() {
	mode := parseFlags()

	fmt.Println(count(os.Stdin, mode))
}

func parseFlags() Mode {
	lines := flag.Bool("l", false, " Count lines")
	bytes := flag.Bool("b", false, " Count bytes")
	flag.Parse()

	if *lines && *bytes {
		log.Fatal("both lines and bytes modes enabled at the same time. Please enable just one at once")
	}

	mode := Words
	if *lines {
		mode = Lines

	} else if *bytes {
		mode = Bytes
	}

	return mode
}

func count(r io.Reader, mode Mode) int {
	scanner := bufio.NewScanner(r)
	if mode == Words {
		scanner.Split(bufio.ScanWords)

	} else if mode == Bytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
