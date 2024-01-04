package game

import (
	"battleships/grid"
	"battleships/player"
)

type Game struct {
	player1 *player.Player
	player2 *player.Player
}

func New(player1, player2 *player.Player) *Game {
	return &Game{player1: player1, player2: player2}
}

func (g Game) FirePlayerOne(row int, col int) grid.ShotResult {
	opponent := g.player2
	return opponent.TakeShot(row, col)
}

func (g Game) FirePlayerTwo(row int, col int) grid.ShotResult {
	return grid.HIT
}