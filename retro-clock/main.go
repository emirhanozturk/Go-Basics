package main

import (
	"fmt"
	"time"
)

func main() {

	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	seperators := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	fmt.Println("\f")
	for {

		fmt.Println("\f")

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

		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
