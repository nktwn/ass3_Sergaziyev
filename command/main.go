package main

import (
	"Sergaziyev_ass3/command/gameengine"
)

func main() {
	engine := gameengine.NewGameEngine()
	player := gameengine.NewPlayer("Player 1")

	// создается игрвоой движок gameEngine который будет управлять игровыми дейтсвиями
	// игроков. выше мы создали нового игрока и он установлен на начальной поиции х 0 и у 0.

	moveCommand := gameengine.NewMoveCommand(player, 3, 4)
	// перемещаем игрока на другую позицию

	attackCommand := gameengine.NewAttackCommand(player)
	// атака

	collectCommand := gameengine.NewCollectCommand(player, "зелье здоровья")
	// позволяет лутать различные предметы

	engine.ExecuteCommand(moveCommand)
	engine.ExecuteCommand(attackCommand)
	engine.ExecuteCommand(collectCommand)
	// каждая команда приводится в действие с помощьюнашего движка

	// отмена последней команды
	engine.Undo()

	// повтор последней отмененной команды
	engine.Redo()
}
