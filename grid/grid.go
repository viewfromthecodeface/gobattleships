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

func NewGrid() *Grid {
	return &Grid{}
}

func isOutOfBounds( row int, col int ) bool {
	return row < 0 || row > 6 || col < 0 || col > 6
}

func (g *Grid) PlaceShip( row int, col int ) error {
	if isOutOfBounds(row, col) {
		return errors.New("ship out of bounds")
	}
		
	if g.isShipAt(row, col) {
		return errors.New("ship already at location")
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

func (g *Grid) IncomingShot( row int, col int ) (ShotResult, error) {
	if isOutOfBounds(row, col) {
		return MISS, errors.New("shot out of bounds")
	}

	if g.isShipAt(row, col) {
		g.sinkShip(row, col)
		return HIT, nil
	}

	return MISS, nil
}

func (g Grid) HasNoShips() bool {
	for _, row := range g.positions {
		for _, position := range row {
			if position != emptySpace {
				return false
			}
		}
	}

	return true
}