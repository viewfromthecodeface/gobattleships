package player

import (
	"battleships/grid"
	"errors"
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
	got, _ := player1.Fire( 2, 3 )

	// Assert
	want := grid.HIT
	if got != want {
		t.Error("did not hit ship")
	}
}

func TestMissesOpponentShip(t *testing.T) {
	// Arrange
	turns := NewTurns()

	player1 := New(turns, "One")
	
	opponent := New(turns, "Opponent")
	opponent.PlaceShip(2, 3)

	// Act
	got, _ := player1.Fire( 0, 0 )

	// Assert
	want := grid.MISS
	if got != want {
		t.Error("did not hit ship")
	}
}

func TestPlayer2TurnBlocked(t *testing.T) {
	// Arrange
	turns := NewTurns()

	New(turns, "One")
	player2 := New(turns, "Two")

	// Act
	_, got := player2.Fire( 0, 0 )

	// Assert
	want := errors.New("not your turn")
	if got.Error() != want.Error() {
		t.Errorf("incorrect error got %v, want %v", got, want)
	}
}

func TestPlayer2GetsATurn(t *testing.T) {
	// Arrange
	turns := NewTurns()

	player1 := New(turns, "One")
	player2 := New(turns, "Two")

	player1.Fire(0, 0) // valid shot (miss) -> player 2 turn

	// Act
	_, err := player2.Fire(0, 0)

	// Assert
	if err != nil {
		t.Error("Unexpected error - player 2 should be allowed to fire")
	}
}
