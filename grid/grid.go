package grid

import (
	"errors"
	"slices"
)

type ship struct {
	row int
	col int
}

type Grid struct {
	shipsPresent []ship
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

	g.addShip(row, col)
	return nil
}

func (g *Grid) addShip(row int, col int) {
	g.shipsPresent = append(g.shipsPresent, 
						    ship{row, col})
}

func (g *Grid) removeShip(target ship) {
	for i, ship := range g.shipsPresent {
		if ship == target {
			g.shipsPresent = append(g.shipsPresent[:i], g.shipsPresent[i+1:]...)
			return
		}
	}
}

func (g Grid) isShipAt( row int, col int ) bool {
	return slices.Contains(g.shipsPresent, ship{row, col})
}

func (g *Grid) sinkShip( row int, col int ) {
	target := ship{row, col}
	g.removeShip(target)
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
	return len(g.shipsPresent) == 0
}