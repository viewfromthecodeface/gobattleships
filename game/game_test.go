package game_test

import (
	"battleships/game"
	"testing"
)

func TestStartWithPlayer1(t *testing.T) {
	// Arrange 
	game := game.New()

	// Act 
	got := game.CurrentPlayer()

	// Assert
	want := game.Player1
	if got != want {
		t.Errorf("got %v, want %v:", got, want)
	}
}

func TestPlayer2FollowsPlayer1Shot(t *testing.T) {
	game := game.New()
	game.TakeShot(2, 3)

	// Act
	got := game.CurrentPlayer()

	want := game.Player2
	if got != want {
		t.Errorf("got %v, want %v:", got, want)
	}
}

func TestPlayer1FollowsPlayer2Shot(t *testing.T) {
	game := game.New()

	game.TakeShot(2, 3) // player 1 shot
	game.TakeShot(2, 3) // player 2 shot

	// Act
	got := game.CurrentPlayer()

	want := game.Player1
	if got != want {
		t.Errorf("got %v, want %v:", got, want)
	}
}

func TestStaysWithPlayer1AfterInvalidShot(t *testing.T) {
	// Arrange 
	game := game.New()
	game.TakeShot(-1, -1) // invalid shot

	// Act 
	got := game.CurrentPlayer()

	// Assert
	want := game.Player1
	if got != want {
		t.Errorf("got %v, want %v:", got, want)
	}
}

