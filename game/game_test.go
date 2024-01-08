package game_test

import (
	"battleships/game"
	"testing"
)

// Stub Text Input source
type StubTextInput struct {
	
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

func TestStartGameDisplaysWelcomeMessage(t *testing.T) {
	// Arrange
	input := StubTextInput{}
	output := &MockTextOutput{}

	game := game.NewGame(input, output)

	// Act
	game.Start()

	// Assert
	output.assertText(t, "Let's play Battleships!")
}