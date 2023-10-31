package gameengine

import (
	"fmt"
)

type GameEngine struct {
	commands       []Command
	currentCommand int
}

func NewGameEngine() *GameEngine {
	return &GameEngine{
		commands:       []Command{},
		currentCommand: -1,
	}
}

func (ge *GameEngine) ExecuteCommand(command Command) {
	ge.currentCommand++
	if ge.currentCommand < len(ge.commands) {
		ge.commands = ge.commands[:ge.currentCommand]
	}
	ge.commands = append(ge.commands, command)
	command.Execute()
}

func (ge *GameEngine) Undo() {
	if ge.currentCommand >= 0 {
		command := ge.commands[ge.currentCommand]
		ge.currentCommand--
		command.Undo()
	} else {
		fmt.Println("Нельзя выполнить отмену.")
	}
}

func (ge *GameEngine) Redo() {
	if ge.currentCommand < len(ge.commands)-1 {
		ge.currentCommand++
		command := ge.commands[ge.currentCommand]
		command.Execute()
	} else {
		fmt.Println("Нельзя выполнить повтор последней отмененной команды.")
	}
}
