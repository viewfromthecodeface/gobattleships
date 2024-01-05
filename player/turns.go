package player

import "battleships/grid"

type Turns struct {
	player2 *Player
}

func NewTurns() *Turns {
	return &Turns{}
}

func (t *Turns) AddPlayer( player *Player ) {
	t.player2 = player
}

func (t Turns) ShootOpponent(row int, col int) grid.ShotResult {
	return t.player2.IncomingShot(row, col)
}