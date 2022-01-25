package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	args := os.Args[1:]
	name, mood := args[0], args[1]

	if len(args) != 2 {
		fmt.Println("[Your name] [positive|negative]")
		return
	}

	moods := [...][3]string{
		{"happy", "awesome", "good"},
		{"sad", "bad", "terrible"},
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(moods[0]))

	var n int

	if mood != "positive" {
		n = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[n][random])
}
