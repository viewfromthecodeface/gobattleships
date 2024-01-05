package player

import (
	"battleships/grid"
)

type Player struct {
	name string
	grid grid.Grid
}

func New( name string ) *Player {
	return &Player{	name: name,
	 				grid: *grid.NewGrid()}
}

func (p Player) GetName() string {
	return p.name
}

func (p *Player) PlaceShip( row int, col int ) error {
	return p.grid.PlaceShip(row, col)
}