package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {

	words := strings.Fields(corpus)
	query := os.Args[1:]

	for _, q := range query {
		q = strings.ToLower(q)
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				break
			}
		}
	}
}
