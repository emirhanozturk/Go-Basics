package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please give a word to search")
		return
	}

	rx := regexp.MustCompile(`[^a-z]+`)

	query := args[0]

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for in.Scan() {
		word := strings.ToLower(in.Text())
		rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}

	if words[query] {
		fmt.Printf("Text contain %q\n", query)
		return
	}
	fmt.Printf("Text does not contain %q\n", query)
}
