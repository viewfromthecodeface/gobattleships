# gobatttleships

This game of battleships is very simple to start:

There are 2 players

Each player has a grid which is 7*7

Each player has 9 Battleships, each of which can occupy only one square on their grid

Each player can place their battleships anywhere on this grid

Players take it in turns to pick any grid square reference

If the player hits a battleship, then it is sunk, and the turn passes to the opponent

If the player misses a battleship then it is called a miss, and the turn passes to the opponent

The player to first sink all their opponent's battleships is the winner

## notes

There are tests in the main_test.go file. The two existing tests are:
1. TestCreateGrid
2. TestPlayerOneTakingShot

There is also a "utility" function getRandomGridSquare that will return a random square on the playing grid using co-ordinates, which is why it returns a slice of int.

There is a single test for our own utility function called TestRandomGridSquare. This test checks that the returned co-ordinates (the slice) are valid.
