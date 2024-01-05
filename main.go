package main

/*
This game of battleships is very simple to start:
There are 2 players
Each player has a grid which is 7*7
Each player has 9 Battleships, each of which can occupy only one square on their grid
Each player can place their battleships anywhere on this grid
Players take it in turns to pick any grid square reference
If the player hits a battleship, then it is sunk, and the turn passes to the opponent
If the player misses a battleship then it is called a miss, and the turn passes to the opponent
The player to first sink all their opponent's battleships is the winner
*/

func main() {
	// not implemented

	// spike
	// turns := Turns{}
	// p1 := player.New(turns, "one")
	// p2 := player.New(turns, "two")

	// p1.Fire(2, 3) // causes p2.Incoming(2, 3)

	// // assume shot hit
	// p1.Fire(0, 0) // error not your turn

	// // p2 shoots
	// p2.Fire(99, 99) // invalid shot
	// p2.Fire(1, 5) // p1.Incoming(1, 5) - miss - back to p1

	// p2.Fire(0,0) - error not your turn
	

	// // I think Im stuck because of the conflict between web driver and console
	// // console is a loop
	// while ( gameInPlay ) {
	// 	print(activePlayerName)
	// 	get input from active player
	// 	fire
	// 	display result
	// 	update active player
	// }

	// but web driver will have an API endpoint for each player on different device
	// there will be a protocol for front end
	// for each player:
	//  while ( my turn )
	//    on click of shot pos
	//    player fire
	//    display result on UI
	//    update active player
	//
	// web API probably has
	// - new game
	// - is it my turn?
	// - fire
	// - result or error UI update
	// - is game over?
	// loop back to "is it my turn"
	// will need some kind of "server push" for it being your turn
	// maybe simple AJAX poll? websocket? long XHR?
}