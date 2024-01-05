package player

import (
	"battleships/grid"
	"errors"
)

type Player struct {
	name string
	grid grid.Grid
	turns *Turns
}

func New( turns *Turns, name string ) *Player {
	player := &Player{	name: name,
	 				grid: *grid.NewGrid(),
					turns: turns,
				}
	
	turns.AddPlayer(player)

	return player
}

func (p Player) GetName() string {
	return p.name
}

func (p *Player) PlaceShip( row int, col int ) error {
	return p.grid.PlaceShip(row, col)
}

func (p *Player) Fire( row int, col int ) (grid.ShotResult, error) {
	if !p.turns.IsActivePlayer(p) {
		return grid.MISS, errors.New("not your turn")
	}

	return p.turns.ShootOpponent(row, col)
}

func (p *Player) IncomingShot( row int, col int ) (grid.ShotResult, error) {
	return p.grid.IncomingShot(row, col)
}