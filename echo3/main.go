package main

import (
	"flag"
	"fmt"
	"strings"
)

var delimiter = flag.String("s", " ", "delimiter char")
var noNewLine = flag.Bool("nonewline", false, "no new line")

func main() {
	flag.Parse()
	// var delimiter string
	// flag.StringVar(&delimiter, "s", " ", "delimiter char")

	// var noNewLine bool
	// flag.BoolVar(&noNewLine, "nonewline", false, "no new line")

	fmt.Printf("%s", strings.Join(flag.Args(), *delimiter))
	if !*noNewLine {
		fmt.Println()
	}
}
