package game

type TextInput interface {

}

type TextOutput interface {
	Show(text string)
}

type Game struct {
	input TextInput
	output TextOutput
}

func NewGame(userInput TextInput, output TextOutput) *Game {
	return &Game{input: userInput, output: output}
}

func (g *Game) Start() {
	g.output.Show( "Let's play Battleships!")
}