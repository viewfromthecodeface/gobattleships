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