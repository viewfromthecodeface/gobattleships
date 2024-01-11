package game

import (
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

type Game struct {
	input TextInput
	output TextOutput
	player1 *player.Player
	player2 *player.Player
}

func NewGame(userInput TextInput, output TextOutput) *Game {
	return &Game{input: userInput, output: output}
}

func (g *Game)FetchShipPosition() (row int, col int) {
	userCommand := g.input.Fetch()

	decimalValue, _ := strconv.Atoi(userCommand)

	row = decimalValue / 10
	col = decimalValue % 10

	return row, col
}

func (g *Game) placeShips(currentPlayer *player.Player) {
	taskOverview := fmt.Sprintf("%s - Place each of your 9 ships", currentPlayer.GetName())
	g.output.Show(taskOverview)

	for ship := 1; ship <= player.MAXIMUM_NUMBER_OF_SHIPS; ship++ {

		prompt := fmt.Sprintf("%s Place Ship %d", currentPlayer.GetName(), ship)
		g.output.Show(prompt)

		row, col := g.FetchShipPosition()
		currentPlayer.PlaceShip(row, col)
	}

	g.output.Show("")
}

func (g *Game) showWelcomeMessage() {
	g.output.Show( "Let's play Battleships!")
	g.output.Show("")
}

func (g *Game) askPlayersToPlaceShips() {
	g.placeShips(g.player1)
	g.placeShips(g.player2)
}

func (g *Game) Play() {
	turns := player.NewTurns()
	g.player1 = player.New(turns, "Player 1")
	g.player2 = player.New(turns, "Player 2")

	g.showWelcomeMessage()
	g.askPlayersToPlaceShips()
}