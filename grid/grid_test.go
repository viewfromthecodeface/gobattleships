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
