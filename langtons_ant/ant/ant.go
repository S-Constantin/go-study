package ant

import (
	"github.com/S-Constantin/go-study/langtons_ant/board"
)

type Direction uint8

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Ant struct {
	y         int
	x         int
	direction Direction
}

func New(Y, X int) Ant {
	return Ant{
		y:         Y,
		x:         X,
		direction: Up,
	}
}

func (a *Ant) turnRight() {
	a.direction = (a.direction + 1) % 4
}

func (a *Ant) turnLeft() {
	a.direction = (a.direction + 3) % 4
}

func (a *Ant) step(width, length int) {

	switch a.direction {
	case Up:
		a.y = (a.y - 1 + length) % length
	case Right:
		a.x = (a.x + 1) % width
	case Down:
		a.y = (a.y + 1) % length
	case Left:
		a.x = (a.x - 1 + width) % width
	}
}

func (a *Ant) Move(b *board.Board) {
	if b.IsBlack(a.y, a.x) {
		a.turnLeft()
	} else {
		a.turnRight()
	}

	b.ChangeColor(a.y, a.x)
	a.step(b.Width(), b.Height())
}

func (a Ant) GetLocation() (int, int) {
	return a.y, a.x
}
