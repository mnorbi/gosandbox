package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var delimiter string
	flag.StringVar(&delimiter, "s", " ", "delimiter char")

	var noNewLine bool
	flag.BoolVar(&noNewLine, "nonewline", false, "no new line")

	flag.Parse()

	fmt.Printf("%s", strings.Join(flag.Args(), delimiter))
	if !noNewLine {
		fmt.Println()
	}
}
