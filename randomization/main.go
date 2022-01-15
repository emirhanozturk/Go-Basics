package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxGuess     = 5
	winMessage1  = "Congrats... YOU WIN"
	winMessage2  = "You are the WINNER!!"
	lostMessage1 = "Oops Sorry... You can win another time..."
	lostMessage2 = "YOU LOST... TRY AGAIN"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Please pick a number.")
		return
	}
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Please pick an int number.")
		return
	}
	if guess < 0 {
		fmt.Printf("Number can not be negative")
		return
	}

	for i := 0; i < maxGuess; i++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			if n == guess && i == 0 {
				fmt.Printf("Congrats... You have found the number first time...")
				return
			}
			m := rand.Intn(2)
			switch m {
			case 0:
				fmt.Printf(winMessage1)
				return
			case 1:
				fmt.Printf(winMessage2)
				return
			}
		}
	}
	m := rand.Intn(2)
	switch m {
	case 0:
		fmt.Printf(lostMessage1)
	case 1:
		fmt.Printf(lostMessage2)
	}

	fmt.Println()
}
