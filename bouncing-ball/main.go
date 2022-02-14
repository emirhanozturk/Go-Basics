package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
)

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20

		ivx, ivy = 5, 2
	)
	var (
		px, py   int
		ppx, ppy int //previous ball position
		vx, vy   = ivx, ivy
		cell     rune
	)
	width, height := screen.Size()
	ballWidth := runewidth.RuneWidth(cellBall)

	width /= ballWidth
	height--

	board := make([][]bool, width)

	for column := range board {
		board[column] = make([]bool, height)
	}
	bufLen := (width*2 + 1) * height

	buf := make([]rune, 0, bufLen)
	screen.Clear()
	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy
		if px <= 0 || px >= width-ivx {
			vx *= -1
		}
		if py <= 0 || py >= height-ivx {
			vy *= -1
		}
		//remove the board
		board[px][py], board[ppx][ppy] = true, false
		ppx, ppy = px, py

		buf = buf[:0]
		//draw the board
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		screen.MoveTopLeft()
		fmt.Println(string(buf))

		time.Sleep(speed)
	}
}
