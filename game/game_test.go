package game_test

import (
	"battleships/game"
	"battleships/grid"
	"battleships/player"
	"testing"
)

// func TestSpike(t *testing.T) {
// 	/* Spiking out how the two player game should look
// 	 * design challenges:
// 	 *  minimise coupling between players
// 	 *  break circular dependency between players
// 	 *  avoid null opponent
// 	 *  decide whose turn is valid
// 	 *  accept commands from player whose turn it is
// 	 *  reject commands from player who is not taking a turn
// 	 *
// 	 * I think this leads to tests like
// 	 * start with player1
// 	 * reject place ship from any player once game started
// 	 * reject take shot from player 2
// 	 * report opponent ship hit from player 1
// 	 * report opponent ship miss
// 	 * stay with player1 following invalid shot
// 	 * move to player2 after valid player1 shot
// 	 * move to player 1 after valid player 2 shot
// 	 * end game after player eins
// 	 *
// 	 * Intuition:
// 	 * we choose a player object, then treat the two equally
// 	 * could have concept of active player, active opponent?
// 	 *
// 	 * pre-game start each player places ships
// 	 * game starts with player 1
// 	 * player 1 take shot command will be accepted
// 	 * if invalid, try again
// 	 * if valid report result
// 	 * switch to other player
// 	 */

// 	 // better design for mediator pattern
// 	 game := game.New()

// 	 player1 := game.NewPlayer1( "one") // link to game
// 	 player1.PlaceShip(1, 2)

// 	 player2 := game.NewPlayer2("two")
// 	 player2.PlaceShip(3, 4)
// 	 player2.PlaceShip(6, 6)

// 	 // design: should the API have a player, then call player.takeshot which delegates to Game?
// 	 // or should it pull stuff out?
// 	 // so player2.TakeShot() would return an err "not your turn"
// 	 // that seems good?
// 	 // implies
// 	 // struct Player
// 	 //  name string
// 	 //  grid Grid
// 	 //  game Game
// 	 // and
// 	 // (p *Player) TakeShot()
// 	 //  if !game.IsActive(p) -> error
// 	 //  else p.TakeShot()

// 	 // Player 1 active
// 	 player1.TakeShot(99, 99) // will stay on player 1, so
// 	 player2.TakeShot(0,0) // will error - not your turn

// 	 player1.TakeShot(3, 4) // hit, not won, move to player 2

// 	 // Player 2 active
// 	 player1.TakeShot(0,0) // error - not your turn
// 	 player2.TakeShot(1,2) // hit - wins game - ends game

// 	 game.IsOver() // true
// 	 game.GetWinner() // => Player 2

// }

func TestReportsPlayer1SunkOpponentShip(t *testing.T) {
	// Arrange
	player1 := player.New("one")

	player2 := player.New("two")
	player2.PlaceShip(2, 3)
	player2.PlaceShip(2, 4)

	game := game.New(player1, player2)

	// Act
	got := game.FirePlayerOne(2, 3)

	// Assert
	want := grid.HIT

	if got != want {
		t.Errorf("ship not hit got %v, want %v", got, want)
	}
}

func TestReportsPlayer1Miss(t *testing.T) {
	// Arrange
	player1 := player.New("one")

	player2 := player.New("two")
	player2.PlaceShip(2, 3)
	player2.PlaceShip(2, 4)

	game := game.New(player1, player2)

	// Act
	got := game.FirePlayerOne(0, 0) // empty location

	// Assert
	want := grid.MISS

	if got != want {
		t.Errorf("ship not missed got %v, want %v", got, want)
	}
}

func TestReportsPlayer2SunkOpponentShip(t *testing.T) {
	// Arrange
	player1 := player.New("one")
	player1.PlaceShip(1, 3)
	player1.PlaceShip(1, 4)
	
	player2 := player.New("two")
	player2.PlaceShip(2, 3)
	player2.PlaceShip(2, 4)

	game := game.New(player1, player2)

	game.FirePlayerOne(2, 3)

	// Act
	got := game.FirePlayerTwo(1, 3)

	// Assert
	want := grid.HIT

	if got != want {
		t.Errorf("ship not hit got %v, want %v", got, want)
	}
}
