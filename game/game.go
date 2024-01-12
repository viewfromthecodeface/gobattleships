package game

import (
	"battleships/grid"
	"battleships/player"
	"fmt"
	"strconv"
)

type TextInput interface {
	Fetch() string
}

type TextOutput interface {
	Show(text string)
}

const MAXIMUM_NUMBER_OF_SHIPS = 9

type Game struct {
	input TextInput
	output TextOutput

	player1 *player.Player
	player2 *player.Player

	shooter *player.Player
	opponent *player.Player
}

func NewGame(userInput TextInput, output TextOutput) *Game {
	return &Game{input: userInput, output: output}
}

func (g *Game) show(text string) {
	g.output.Show(text)
}

func (g *Game)FetchShipPosition() (row int, col int) {
	userCommand := g.input.Fetch()
	decimalValue, _ := strconv.Atoi(userCommand)

	row = decimalValue / 10
	col = decimalValue % 10

	return row, col
}

func (g *Game) placeShips(currentPlayer *player.Player) {
	taskOverview := fmt.Sprintf("%s - Place each of your 9 ships", currentPlayer.Name)
	g.show(taskOverview)

	for ship := 1; ship <= MAXIMUM_NUMBER_OF_SHIPS; ship++ {

		prompt := fmt.Sprintf("%s Place Ship %d", currentPlayer.Name, ship)
		g.output.Show(prompt)

		row, col := g.FetchShipPosition()
		currentPlayer.PlaceShip(row, col)
	}

	g.output.Show("")
}

func (g *Game) showWelcomeMessage() {
	g.show( "Let's play Battleships!")
	g.show("")
}

func (g *Game) askPlayersToPlaceShips() {
	g.placeShips(g.player1)
	g.placeShips(g.player2)
}

func (g *Game) otherPlayerShoots() {
	previousShooter := g.shooter

	g.shooter = g.opponent
	g.opponent = previousShooter
}

func (g *Game) fire() grid.ShotResult {
	for {
		prompt := fmt.Sprintf("%s - enter position 00-66 of your shot", 
								g.shooter.Name)
		g.show(prompt)

		row, col := g.FetchShipPosition()
		shotResult, err := g.opponent.IncomingShot(row, col)

		if err == nil {
			return shotResult
		} 
			
		g.show("invalid position, try again")
	}
}

func (g *Game) Play() {
	g.player1 = player.NewPlayer("Player 1", MAXIMUM_NUMBER_OF_SHIPS)
	g.player2 = player.NewPlayer("Player 2", MAXIMUM_NUMBER_OF_SHIPS)

	g.showWelcomeMessage()
	g.askPlayersToPlaceShips()
	
	g.show("Let's Play!")

	g.shooter = g.player1
	g.opponent = g.player2

	gameOver := false
	for !gameOver {
		shotResult := g.fire()
		
		if shotResult == grid.HIT {
			g.show("HIT! Enemy ship sunk")

			if g.opponent.AllShipsSunk() {
				gameOver = true
			} else {
				g.otherPlayerShoots()
			}
		} else {
			g.show("shot missed")
			g.otherPlayerShoots()
		}
	}

	g.show(g.shooter.Name + " WINS!")
}