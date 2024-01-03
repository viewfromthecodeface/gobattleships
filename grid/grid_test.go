package grid

import (
	"testing"
)

func TestCreatesEmptyGrid(t *testing.T){
	// Arrange
	// No Action

	// Act
	grid := NewGrid()

	// Assert
	got := grid.HasNoShips()
	want := true

	if got != want {
		t.Errorf("grid not empty")
	}
}

func TestPlaceShip(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	grid.PlaceShip(1, 2)

	// Assert
	if !grid.isShipAt(1, 2) {
		t.Error("Ship not placed")
	}
}

func TestReportsShipHit(t *testing.T) {
	// Arrange
	grid := NewGrid()
	grid.PlaceShip(1, 2)

	// Act
	got := grid.TakeShot(1, 2)

	// Arrange
	want := HIT

	if got != want {
		t.Errorf("Did not report hit. got %v, want %v", got, want)
	}
}

func TestReportsMiss(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act 
	got := grid.TakeShot(0, 0)

	// Arrange
	want := MISS

	if got != want {
		t.Errorf("Did not report miss. got %v, want %v", got, want)
	}
}

func TestReportsMissForShotAtAlreadySunkShip(t *testing.T) {
	// Arrange - place then sink one ship
	grid := NewGrid()
	
	grid.PlaceShip(1, 2)
	grid.TakeShot(1, 2)	// sink ship

	// Act 
	got := grid.TakeShot(1, 2)

	// Arrange
	want := MISS

	if got != want {
		t.Errorf("Did not report miss. got %v, want %v", got, want)
	}
}

func TestPlacesShipTopLeft(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	err := grid.PlaceShip(0, 0)

	if err != nil {
		t.Error()
	}
}

func TestPlacesShipTopRight(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	err := grid.PlaceShip(6, 0)

	if err != nil {
		t.Error()
	}
}

func TestPlacesShipBottomLeft(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	err := grid.PlaceShip(0, 6)

	if err != nil {
		t.Error()
	}
}

func TestPlacesShipBottomRight(t *testing.T) {
	// Arrange
	grid := NewGrid()

	// Act
	err := grid.PlaceShip(6, 6)

	if err != nil {
		t.Error()
	}
}

func TestReportsShipsOnGrid(t *testing.T) {
	// Arrange
	grid := NewGrid()

	grid.PlaceShip(1, 2)

	// Act
	got := grid.HasNoShips()

	// Assert
	want := false

	if got != want {
		t.Error()
	} 
}