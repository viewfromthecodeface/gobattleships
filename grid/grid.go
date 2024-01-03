package grid

type Grid struct {
	positions [7][7]string
}

const emptySpace string = ""

func NewGrid() *Grid {
	return &Grid{}
}

func (g Grid) isEmpty(row int, col int) bool {
	return g.positions[row][col] == emptySpace
}

func (g Grid) IsEmpty() bool {
	for _, row := range g.positions {
		for _, position := range row {
			if position != emptySpace {
				return false
			}
		}
	}

	return true
}