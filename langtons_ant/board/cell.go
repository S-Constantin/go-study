package board

type Cell struct {
	black bool
}

func (c *Cell) setBlack() {
	c.black = true
}

func (c *Cell) setWhite() {
	c.black = false
}
