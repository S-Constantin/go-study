package main

import (
	"time"

	"github.com/S-Constantin/go-study/langtons_ant/ant"
	"github.com/S-Constantin/go-study/langtons_ant/board"
	"github.com/S-Constantin/go-study/langtons_ant/read"
)

var (
	width      int
	height     int
	iterations int
	delay      time.Duration
)

func main() {

	read.New(&width, &height, &iterations, &delay)
	
	gameBoard := board.New(height, width)
	langtonAnt := ant.New(height/2, width/2)

	for i := 1; i <= iterations; i++ {
		gameBoard.Show(langtonAnt.GetLocation())

		langtonAnt.Move(gameBoard)

		time.Sleep(delay)
	}
}
