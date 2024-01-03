package grid

import "testing"

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
