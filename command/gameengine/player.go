package gameengine

import "fmt"

type Player struct {
	name      string
	x, y      int
	inventory []string
}

func NewPlayer(name string) *Player {
	return &Player{
		name:      name,
		x:         0,
		y:         0,
		inventory: []string{},
	}
}

func (p *Player) MoveTo(x, y int) {
	p.x = x
	p.y = y
	fmt.Printf("%s переместился в позицию (%d, %d)\n", p.name, x, y)
}

func (p *Player) Attack() {
	fmt.Printf("%s атакует противника\n", p.name)
}

func (p *Player) CollectItem(itemName string) {
	p.inventory = append(p.inventory, itemName)
	fmt.Printf("%s подобрал %s\n", p.name, itemName)
}

func (p *Player) DropItem(itemName string) {
	for i, item := range p.inventory {
		if item == itemName {
			p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
			fmt.Printf("%s выбросил %s\n", p.name, itemName)
			return
		}
	}
}
