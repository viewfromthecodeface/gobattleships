package game_test

import (
	"battleships/game"
	"battleships/grid"
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

func TestStaysWithPlayer2AfterInvalidShot(t *testing.T) {
	// Arrange 
	game := game.New()
	game.TakeShot(2, 3)   // Player 1 valid shot
	game.TakeShot(-1, -1) // Player 2 invalid shot

	// Act 
	got := game.CurrentPlayer()

	// Assert
	want := game.Player2
	if got != want {
		t.Errorf("got %v, want %v:", got, want)
	}
}

func TestReportsHitPlayer1 (t *testing.T) {
	game := game.New()
	game.Player2.PlaceShip(2, 3)

	got, _ := game.TakeShot(2, 3) // Player 1 hit

	want := grid.HIT
	if got != want {
		t.Errorf("got %v, want %v:", got, want)
	}
}

func TestReportsPlayerShotResult(t *testing.T) {
	// Arrange 
	game := game.New()

	// Act 
	got, _ := game.TakeShot(2, 3) // player 1 miss

	// Assert
	want := grid.MISS
	if got != want {
		t.Errorf("got %v, want %v:", got, want)
	}
}


func TestReportsPlayerInvalidShot(t *testing.T) {
	// Arrange 
	game := game.New()

	// Act 
	_, got := game.TakeShot(-1, -1) // player 1 miss

	// Assert
	if got == nil {
		t.Error("error not reported")
	}
}

func TestReportsPlayer2Wins(t *testing.T) {
	// Arrange 
	game := game.New()
	game.Player1.PlaceShip(0, 0)
	game.Player2.PlaceShip(6, 6)
	game.TakeShot(1, 1) // player 1 miss
	game.TakeShot(0, 0) // player 2 hit and win
	
	// Act 
	got := game.Winner

	// Assert
	want := game.Player2
	if got != want {
		t.Errorf("got %v, want %v:", got, want)
	}
}