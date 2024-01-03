package grid

type Grid struct {
	positions [7][7]string
}

const emptySpace string = ""

func NewGrid() *Grid {
	return &Grid{}
}