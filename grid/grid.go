package grid

import "errors"

type Grid struct {
	positions [7][7]string
}

const (
	emptySpace string = ""
	shipToken string = "ship"
)

type ShotResult int
const (
	MISS = ShotResult (iota)
	HIT
)

const (
	MAXIMUM_NUMBER_OF_SHIPS int = 9
)

func NewGrid() *Grid {
	return &Grid{}
}

func (g Grid) reachedMaximumShips() bool { 
	total := 0

	for _, row := range g.positions {
		for _, position := range row {
			if position == shipToken {
				total++
			}
		}
	}

	return total == MAXIMUM_NUMBER_OF_SHIPS
}

func (g *Grid) PlaceShip( row int, col int ) error {
	if row < 0 {
		return errors.New("ship out of bounds")
	}
		if g.isShipAt(row, col) {
		return errors.New("ship already at location")
	}

	if g.reachedMaximumShips() {
		return errors.New("cannot place ship - limit reached")
	}

	g.positions[row][col] = shipToken
	return nil
}

func (g Grid) isShipAt( row int, col int ) bool {
	return g.positions[row][col] == shipToken
}

func (g *Grid) sinkShip( row int, col int ) {
	g.positions[row][col] = emptySpace
}

func (g *Grid) TakeShot( row int, col int ) ShotResult {
	if g.isShipAt(row, col) {
		g.sinkShip(row, col)
		return HIT
	}

	return MISS
}