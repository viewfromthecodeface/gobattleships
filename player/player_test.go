package player

import "testing"

func TestReportsPlayerName(t *testing.T) {
	// Arrange
	player := NewPlayer("Action Dan")

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
	player := NewPlayer("One")
	
	// Act
	err := player.PlaceShip(1, 2)

	// Assert
	if err != nil {
		t.Errorf("error placing ship: %v", err)
	}
}

/* 
  TODO next tests: 
  - Places ships on own grid (implies own grid)
  - Take shot at opponent (implies link to opponent grid)
  - PlayerHasWonWhenAllShipsSunk
*/