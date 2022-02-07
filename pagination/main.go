package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {

	items := []string{
		"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteroids", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima",
	}

	const pageSize = 4

	l := len(items)

	for from := 0; from < l; from += pageSize {
		to := from + pageSize

		if to > l {
			to = l
		}
		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from/pageSize)+1)
		s.Show(head, currentPage)
	}

}
