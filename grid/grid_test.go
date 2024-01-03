package grid

import (
	"errors"
	"testing"
)

func TestCreatesEmptyGrid(t *testing.T){
	// Arrange
	// No Action

	// Act
	grid := NewGrid()

	// Assert
	got := isGridEmpty(grid)
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

func isGridEmpty(g *Grid) bool {
	for _, row := range g.positions {
		for _, position := range row {
			if position != emptySpace {
				return false
			}
		}
	}

	return true
}
