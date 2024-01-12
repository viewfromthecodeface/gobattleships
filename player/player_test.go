package player

import (
	"battleships/grid"
	"errors"
	"testing"
)

const NUMBER_OF_SHIPS = 9

func TestReportsPlayerName(t *testing.T) {
	// Arrange
	player := NewPlayer("Action Dan", NUMBER_OF_SHIPS)

	// Act
	got := player.Name

	// Assert
	want := "Action Dan"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPlacesShipWithoutError(t *testing.T) {
	// Arrange
	player := NewPlayer("One", NUMBER_OF_SHIPS)
	
	// Act
	err := player.PlaceShip(1, 2)

	// Assert
	if err != nil {
		t.Errorf("error placing ship: %v", err)
	}
}

func TestReportsErrorTooManyShipsPlaced(t *testing.T) {
	// Arrange - place maximum limit of 9 ships
	player := NewPlayer("One", 9)

	player.PlaceShip(0, 0)
	player.PlaceShip(0, 1)
	player.PlaceShip(0, 2)

	player.PlaceShip(1, 0)
	player.PlaceShip(1, 1)
	player.PlaceShip(1, 2)

	player.PlaceShip(2, 0)
	player.PlaceShip(2, 1)
	player.PlaceShip(2, 2)

	// Act
	got := player.PlaceShip(3, 0)

	// Assert
	want := errors.New("cannot place ship - limit reached")

	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReportsHit(t *testing.T) {
	player := NewPlayer("One", 9)
	player.PlaceShip(2, 3)

	shotResult, _ := player.IncomingShot(2, 3)

	if shotResult != grid.HIT {
		t.Error("did not report hit")
	}
}