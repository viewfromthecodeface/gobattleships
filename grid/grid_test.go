package grid

import "testing"

func TestCreatesEmptyGrid(t *testing.T){
	// Arrange
	grid := NewGrid()

	// Act
	got := grid.IsEmpty()

	// Assert
	want := true
	if got != want {
		t.Errorf("grid not empty")
	}
}
