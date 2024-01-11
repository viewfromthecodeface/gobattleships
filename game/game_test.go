package game_test

import (
	"battleships/game"
	"testing"
)

// Stub Text Input source
type StubTextInput struct {
	inputs []string
	current int
}

func (t *StubTextInput) Fetch() string {
	nextInput := t.inputs[t.current]
	t.current++
	return nextInput 
}

type MockTextOutput struct {
	textHistory []string
}

func (o *MockTextOutput) Show(text string) {
	o.textHistory = append(o.textHistory, text)
}

func (o *MockTextOutput) assertText(t *testing.T, want string) {
	got := o.textHistory[0]

	if got != want {
		t.Errorf("Missing output text: got %v, want %v", got, want)
	}
}

func (o *MockTextOutput) assertAllText(t *testing.T, allWanted []string) {
	actualIndex :=0

	for _, want := range allWanted {
		if actualIndex >= len(o.textHistory) {
			t.Errorf("Missing output - not enough text sent. got %v", o.textHistory)
			return
		}
		
		got := o.textHistory[actualIndex]
		actualIndex++

		if got != want {
			t.Errorf("Missing output text: got %v, want %v", got, want)
		} 
	}
}

func TestPlayer1WinsRepresentativeGame(t *testing.T) {
	// Arrange
	allInputBothPlayers := []string{
		// Player 1 place ship x 9
		"10", "11", "12", "13", "14","15", "16", "20", "21",

		// Player 2 place ship x 9
		"60", "61", "62", "63", "64","65", "66", "50", "51",

		// Player 1 first shot - invalid coordinates
		"99",

		// Player 1 then Player 2 shots
		"00", "00", // miss miss
		"60", "00", // hit miss
		"61", "00", // hit miss
		"62", "00", // hit miss
		"63", "00", // hit miss
		"64", "00", // hit miss
		"65", "00", // hit miss
		"50", "00", // hit miss
		"51", "00", // hit miss
		"52", "00", // winning-hit miss
	}

	input := &StubTextInput{ inputs: allInputBothPlayers }
	output := &MockTextOutput{}

	game := game.NewGame(input, output)

	// Act
	game.Play()

	// Assert
	want := []string{
		"Let's play Battleships!",
		"",
		"Player 1 - Place each of your 9 ships",
		"Player 1 Place Ship 1",
		"Player 1 Place Ship 2",
		"Player 1 Place Ship 3",
		"Player 1 Place Ship 4",
		"Player 1 Place Ship 5",
		"Player 1 Place Ship 6",
		"Player 1 Place Ship 7",
		"Player 1 Place Ship 8",
		"Player 1 Place Ship 9",
		"",
		"Player 2 - Place each of your 9 ships",
		"Player 2 Place Ship 1",
		"Player 2 Place Ship 2",
		"Player 2 Place Ship 3",
		"Player 2 Place Ship 4",
		"Player 2 Place Ship 5",
		"Player 2 Place Ship 6",
		"Player 2 Place Ship 7",
		"Player 2 Place Ship 8",
		"Player 2 Place Ship 9",
		"",
		"Let's Play!",
		"Player 1 - enter position 00-66 of your shot",
		"invalid position, try again",
	}

	output.assertAllText(t, want)
}

func TestFetchesRowShipPosition(t *testing.T) {
	// Arrange
	input := &StubTextInput{ inputs: []string{"23"} }
	game := game.NewGame(input, nil)

	row, _ := game.FetchShipPosition()

	if row != 2 {
		t.Errorf("wrong row got %v, want 2", row)
	}
}

func TestFetchesColumnShipPosition(t *testing.T) {
	// Arrange
	input := &StubTextInput{ inputs: []string{"23"} }
	game := game.NewGame(input, nil)

	_, col := game.FetchShipPosition()

	if col != 3 {
		t.Errorf("wrong col got %v, want 3", col)
	}
}