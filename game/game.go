package game

import "battleships/player"

type Game struct {
	player1 *player.Player
	//player2 *player.Player
	//activePlayer *player.Player
}

func New(player1 *player.Player) *Game {
	return &Game{player1: player1}
}

func (g Game) GetActivePlayer() *player.Player {
	return g.player1
}
