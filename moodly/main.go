package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	args := os.Args[1:]
	name := args[0]

	if len(args) != 1 {
		fmt.Println("[Your name]")
		return
	}

	moods := [...]string{
		"sad",
		"happy",
		"terrible",
		"awesome",
		"bad",
		"good",
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(moods))

	fmt.Printf("%s feels %s", name, moods[random])
}
