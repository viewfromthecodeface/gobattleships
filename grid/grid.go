package grid

type Grid struct {
	positions [7][7]string
}

const (
	emptySpace = ""
	shipToken string = "ship"
)

type ShotResult int
const (
	MISS = ShotResult (iota)
	HIT
)

func NewGrid() *Grid {
	return &Grid{}
}

func (g *Grid) PlaceShip( row int, col int ) {
	g.positions[row][col] = shipToken
}

func (g Grid) isShipAt( row int, col int ) bool {
	return g.positions[row][col] == shipToken
}

func (g *Grid) TakeShot( row int, col int ) ShotResult {
	if g.positions[row][col] == shipToken {
		return HIT
	}

	return MISS
}