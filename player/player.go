package player

type Player struct {
	name string
}

func NewPlayer( name string ) *Player {
	return &Player{name: name}
}

func (p Player) GetName() string {
	return p.name
}