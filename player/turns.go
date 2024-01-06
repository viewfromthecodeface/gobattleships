package player

import (
	"battleships/grid"
)

type Turns struct {
	player1 *Player
	player2 *Player

	activePlayer *Player
	opponent *Player

	winner *Player
}

func NewTurns() *Turns {
	return &Turns{}
}

func (t *Turns) AddPlayer( player *Player ) {
	if t.player1 == nil {
		t.player1 = player
		t.activePlayer = t.player1
		return
	}

	t.player2 = player
	t.opponent = t.player2
}

func (t *Turns) swapPlayers() {
	if t.activePlayer == t.player1 {
		t.activePlayer = t.player2
		t.opponent = t.player1
		return 
	} 

	t.activePlayer = t.player1
	t.opponent = t.player2
}

func (t *Turns) updateTurn(err error) {
	wasValidShot := err == nil

	if wasValidShot {
		t.swapPlayers()
	}
}

func (t *Turns) ShootOpponent(row int, col int) (grid.ShotResult, error) {
	result, err := t.opponent.IncomingShot(row, col)

	t.updateTurn(err)

	return result, err
}

func (t Turns) IsActivePlayer(player *Player) bool {
	return t.activePlayer == player
}

func (t Turns) IsGameWon() bool {
	return t.winner != nil
}

func (t *Turns) RecordWin() {
	t.winner = t.activePlayer
}