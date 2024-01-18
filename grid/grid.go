package grid

import "errors"

type Grid struct {
	shipsPresent [49]bool // true if ship present
}

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

	g.shipsPresent[toIndex(row, col)] = true
	return nil
}

func toIndex( row int, col int ) int {
	return (row * 7) + col
}

func (g Grid) isShipAt( row int, col int ) bool {
	return g.shipsPresent[toIndex(row, col)]
}

func (g *Grid) sinkShip( row int, col int ) {
	g.shipsPresent[toIndex(row,col)] = false
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
	for _, isShipPresent := range g.shipsPresent {
		if isShipPresent {
			return false
		}
	}

	return true
}