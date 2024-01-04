package game

import "battleships/player"

type Game struct {
	player1 *player.Player
	player2 *player.Player
}

func New(p1 *player.Player, p2 *player.Player) *Game {
	return &Game{player1: p1, player2: p2}
}