package model

import "fmt"

type Player struct {
	ID    string
	State PlayerState
}

func (p *Player) String() string {
	return fmt.Sprintf("ID: %s, State: %v", p.ID, p.State)
}

type PlayerState struct {
	X      float64
	Y      float64
	DeltaX float64
	DeltaY float64
}
