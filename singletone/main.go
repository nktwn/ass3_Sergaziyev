package main

import (
	"Sergaziyev_ass3/singletone/settings"
	"fmt"
)

// данный проект реализует приложение для игр который позволяет различным игрокам менять
// базовые настройки игры. благодарня паттерну singletone, все игроки являющийся структурами, работают
// с одним и тем же наборем настроек и гарантирует, что изменения внесенные одним игроком
// будут отображься в аккауенте и других игроков.

func main() {
	player1 := settings.NewPlayer("Player 1")
	// создаются два игрока, и для каждого из них вызывается метод ChangeSettings,
	// чтобы изменить настройки игры. после изменения настроек все игроки
	// могут просматривать обновленные настройки.
	gm := settings.NewGameManager()
	initialDifficulty := gm.GetDifficulty()
	initialHints := gm.GetHints()
	initialControl := gm.GetControl()
	fmt.Println("Initial Game Settings:")
	fmt.Printf("Difficulty: %s, Hints: %s, Corntol: %s\n", initialDifficulty, initialHints, initialControl)
	fmt.Println()

	player1.ChangeSettings("Hard", "Off", "Gamepad")

	gm2 := settings.NewGameManager()
	updatedDifficulty := gm2.GetDifficulty()
	updatedHints := gm2.GetHints()
	updatedControl := gm2.GetControl()
	fmt.Println("Updated game settings (Player 2):")
	fmt.Printf("Difficult: %s, Hints: %s, Control: %s\n", updatedDifficulty, updatedHints, updatedControl)
}
