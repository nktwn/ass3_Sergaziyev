package settings

import (
	"fmt"
)

// создание новых игроков
type Player struct {
	name string
}

func NewPlayer(name string) *Player {
	return &Player{name}
}

func (p *Player) ChangeSettings(difficulty string, hints string, control string) {
	gm := NewGameManager()
	gm.SetDifficulty(difficulty)
	gm.SetHints(hints)
	gm.SetControl(control)
	fmt.Printf("%s changed settings:\n", p.name)
	fmt.Printf("Difficulty: %s, Hints: %s, Control: %s\n", difficulty, hints, control)
}
