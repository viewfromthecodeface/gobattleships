package player

import (
	"battleships/grid"
	"errors"
)

type Player struct {
	Name string
	grid grid.Grid
	numberOfShipsToPlace int
}

func NewPlayer( name string, numberOfShips int ) *Player {
	return &Player{	Name: name,
	 				grid: *grid.NewGrid(),
					numberOfShipsToPlace: numberOfShips,
				}
}

func (p *Player) PlaceShip( row int, col int ) error {
	if p.numberOfShipsToPlace == 0 {
		return errors.New("cannot place ship - limit reached")
	}

	p.numberOfShipsToPlace--
	return p.grid.PlaceShip(row, col)
}

func (p *Player) IncomingShot( row int, col int ) (grid.ShotResult, error) {
	return p.grid.IncomingShot(row, col)
}

func (p *Player) AllShipsSunk() bool {
	return p.grid.HasNoShips()
} 
