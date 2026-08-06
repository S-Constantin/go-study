package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/S-Constantin/go-study/langtons_ant/ant"
	"github.com/S-Constantin/go-study/langtons_ant/board"
)

var (
	width      int
	height     int
	iterations int
	delay      time.Duration
)

func main() {
	flag.IntVar(&width, "width", 20, "ширина игрового поля")
	flag.IntVar(&height, "height", 20, "высота игрового поля")
	flag.IntVar(&iterations, "iterations", 500, "количество итераций")
	flag.DurationVar(&delay, "delay", 100*time.Millisecond, "задержка между кадрами")

	flag.Parse()

	if width <= 0 {
		fmt.Fprintf(os.Stderr, "ширина должна быть больше нуля")
		os.Exit(1)
	}
	if height <= 0 {
		fmt.Fprintf(os.Stderr, "высота должна быть больше нуля")
		os.Exit(1)
	}
	if iterations <= 0 {
		fmt.Fprintf(os.Stderr, "кол-во иетраций должно быть больше нуля")
		os.Exit(1)
	}
	if delay <= 0 {
		fmt.Fprintf(os.Stderr, "задержка должна быть больше нуля")
		os.Exit(1)
	}

	gameBoard := board.New(height, width)
	langtonAnt := ant.New(height/2, width/2)

	for i := 1; i <= iterations; i++ {
		gameBoard.Show(langtonAnt.GetLocation())

		langtonAnt.Move(gameBoard)

		time.Sleep(delay)
	}
}
