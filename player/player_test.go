package player

import (
	"battleships/grid"
	"testing"
)

func TestReportsPlayerName(t *testing.T) {
	// Arrange
	turns := NewTurns()
	player := New(turns, "Action Dan")

	// Act
	got := player.GetName()

	// Assert
	want := "Action Dan"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPlaceShipOnOwnGrid(t *testing.T) {
	// Arrange
	turns := NewTurns()
	player := New(turns, "One")
	
	// Act
	err := player.PlaceShip(1, 2)

	// Assert
	if err != nil {
		t.Errorf("error placing ship: %v", err)
	}
}

func TestHitsOpponentShip(t *testing.T) {
	// Arrange
	turns := NewTurns()

	player1 := New(turns, "One")
	
	opponent := New(turns, "Opponent")
	opponent.PlaceShip(2, 3)

	// Act
	got := player1.Fire( 2, 3 )

	// Assert
	want := grid.HIT
	if got != want {
		t.Error("did not hit ship")
	}
}