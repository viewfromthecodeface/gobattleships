package player

import (
	"battleships/grid"
	"errors"
)

type Player struct {
	name string
	grid grid.Grid
	turns *Turns
	numberOfShips int
}

const (
	MAXIMUM_NUMBER_OF_SHIPS int = 9
)


func New( turns *Turns, name string ) *Player {
	player := &Player{	name: name,
	 				grid: *grid.NewGrid(),
					turns: turns,
					numberOfShips: 0,
				}
	
	turns.AddPlayer(player)

	return player
}

func (p Player) GetName() string {
	return p.name
}

func (p *Player) PlaceShip( row int, col int ) error {
	if p.numberOfShips == MAXIMUM_NUMBER_OF_SHIPS {
		return errors.New("cannot place ship - limit reached")
	}

	p.numberOfShips++

	return p.grid.PlaceShip(row, col)
}

func (p *Player) Fire( row int, col int ) (grid.ShotResult, error) {
	if !p.turns.IsActivePlayer(p) {
		return grid.MISS, errors.New("not your turn")
	}

	return p.turns.ShootOpponent(row, col)
}

func (p *Player) updateWinStatus( ) {
	if p.grid.HasNoShips() {
		p.turns.RecordWin()
	}
}

func (p *Player) IncomingShot( row int, col int ) (grid.ShotResult, error) {
	shotResult, err := p.grid.IncomingShot(row, col)

	if shotResult == grid.HIT {
		p.updateWinStatus()
	}

	return shotResult, err
}