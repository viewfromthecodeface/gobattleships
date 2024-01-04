package game

import "battleships/player"

type Game struct {
	player1 *player.Player
	player2 *player.Player
}

func New() *Game {
	return &Game{}
}

func (g *Game) CreatePlayer1( name string ) *player.Player {
	g.player1 = player.NewPlayer(name)
	return g.player1
}

func (g *Game) CreatePlayer2( name string ) *player.Player {
	g.player2 = player.NewPlayer(name)
	return g.player2
}