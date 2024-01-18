package game

import (
	"battleships/grid"
	"battleships/player"
)

const MAXIMUM_NUMBER_OF_SHIPS = 9

type Game struct {
	Player1 *player.Player
	Player2 *player.Player
	Winner *player.Player

	shooter *player.Player
	opponent *player.Player
}

func New() *Game {
	g := &Game{}

	g.Player1 = player.NewPlayer("Player 1", MAXIMUM_NUMBER_OF_SHIPS)
	g.Player2 = player.NewPlayer("Player 2", MAXIMUM_NUMBER_OF_SHIPS)
	
	g.shooter = g.Player1
	g.opponent = g.Player2
	
	return g
}

func (g *Game) CurrentPlayer() *player.Player {
	return g.shooter
}

func (g *Game) nextTurn() {
	// swap shooter <-> opponent
	g.shooter, g.opponent = g.opponent, g.shooter
}

func (g *Game) TakeShot(row int, col int) (grid.ShotResult, error) {
	shotResult, err := g.opponent.IncomingShot(row, col)

	isValidShot := err == nil
	if !isValidShot {
		return grid.MISS, err
	}

	if shotResult == grid.HIT {
		if g.opponent.AllShipsSunk() {
			g.Winner = g.shooter
		}
	}

	g.nextTurn()
	return shotResult, err
}