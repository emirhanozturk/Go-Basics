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

		fmt.Printf("Hour:%d,min:%d,second:%d\n", hour, minute, second)

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			seperators,
			digits[minute/10], digits[minute%10],
			seperators,
			digits[second/10], digits[second%10],
		}
		alarmed := second%10 == 0

		for line := range clock[0] {
			if alarmed {
				clock = alarm
			}
			for index, digit := range clock {
				next := clock[index][line]
				if digit == seperators && second%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
