package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()

	for {

		screen.MoveTopLeft()

		now := time.Now()
		hour := now.Hour()
		minute := now.Minute()
		second := now.Second()
		ssec := now.Nanosecond() / 1e8

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			seperators,
			digits[minute/10], digits[minute%10],
			seperators,
			digits[second/10], digits[second%10],
			dot,
			digits[ssec],
		}
		for line := range clock[0] {

			for index, digit := range clock {
				next := clock[index][line]
				if (digit == seperators || digit == dot) && second%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		const splitSecond = time.Second / 10
		time.Sleep(splitSecond)
	}
}
