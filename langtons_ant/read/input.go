package read

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func isCorrect(width, height, iterations int, delay time.Duration) {
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
}

func New(width, height, iterations *int, delay *time.Duration) {
	flag.IntVar(width, "width", 20, "ширина игрового поля")
	flag.IntVar(height, "height", 20, "высота игрового поля")
	flag.IntVar(iterations, "iterations", 500, "количество итераций")
	flag.DurationVar(delay, "delay", 100*time.Millisecond, "задержка между кадрами")

	flag.Parse()

	isCorrect(*width, *height, *iterations, *delay)
}
