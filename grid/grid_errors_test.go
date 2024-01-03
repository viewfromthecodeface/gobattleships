package grid

import (
	"errors"
	"testing"
)

func TestCannotPlaceShipOnTopOfAnother(t *testing.T) {
	// Arrange 
	grid := NewGrid()
	
	grid.PlaceShip(1, 2)

	// Act 
	got := grid.PlaceShip(1, 2)

	// Arrange
	want := errors.New("ship already at location")

	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReportsErrorTooManyShipsPlaced(t *testing.T) {
	// Arrange - place maximum limit of 9 ships
	grid := NewGrid()
	
	grid.PlaceShip(0, 0)
	grid.PlaceShip(0, 1)
	grid.PlaceShip(0, 2)

	grid.PlaceShip(1, 0)
	grid.PlaceShip(1, 1)
	grid.PlaceShip(1, 2)

	grid.PlaceShip(2, 0)
	grid.PlaceShip(2, 1)
	grid.PlaceShip(2, 2)

	// Act
	got := grid.PlaceShip(3, 0)

	// Assert
	want := errors.New("cannot place ship - limit reached")

	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReportsErrorPlacingShipOutsideLeft(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	got := grid.PlaceShip(-1, 0)

	// Assert
	assertErrorText(t, got, "ship out of bounds")
}

func TestReportsErrorPlacingShipOutsideRight(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	got := grid.PlaceShip(7, 0)

	// Assert
	assertErrorText(t, got, "ship out of bounds")
}

func TestReportsErrorPlacingShipOutsideTop(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	got := grid.PlaceShip(0, -1)

	// Assert
	assertErrorText(t, got, "ship out of bounds")
}
func TestReportsErrorPlacingShipOutsideBottom(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	got := grid.PlaceShip(0, 7)

	// Assert
	assertErrorText(t, got, "ship out of bounds")
}

func assertErrorText(t *testing.T, got error, wantedText string) {
	want := errors.New(wantedText)

	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

