package board

import (
	"fmt"
)

type Board struct {
	height int
	width  int
	field  [][]Cell
}

func (b Board) Width() int {
	return b.width
}

func (b Board) Height() int {
	return b.height
}

func (b Board) IsBlack(y, x int) bool {
	return b.field[y][x].black
}

func (b *Board) ChangeColor(y, x int) {
	if b.IsBlack(y, x) {
		b.field[y][x].setWhite()
	} else {
		b.field[y][x].setBlack()
	}
}

func New(height, width int) *Board {

	block := make([]Cell, height*width)
	for i := range block {
		block[i].setWhite()
	}

	new_board := make([][]Cell, height)
	for i := range new_board {
		new_board[i] = block[width*i : (i+1)*width]
	}

	return &Board{
		height: height,
		width:  width,
		field:  new_board,
	}
}

func (b Board) Show(y, x int) {
	for Oy := range b.field {
		for Ox, square := range b.field[Oy] {
			if Ox != x || Oy != y {

				if square.black {
					fmt.Printf(" %c ", '○')
				} else {
					fmt.Printf(" %c ", '●')
				}

			} else {
				fmt.Print(" M ")
			}
		}
		fmt.Println()
	}

	fmt.Printf("Координаты муравья Лэнгтона - Ox: %v, Oy: %v\n\n", x, y)
}
