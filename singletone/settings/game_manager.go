package settings

import (
	"sync"
)

type GameSettings struct {
	Difficulty string
	Hints      string
	Control    string
}

// создается структура GameManager, которая позволяет пользователям
// получать и изменять настройки игры. этот менеджер действует как Singleton,
// то есть создает экземпляр  при первом обращении к нему и затем предоставляет
// данный экземпляр для каждого последующего обращения.

type GameManager struct {
	settings *GameSettings
	mu       sync.Mutex
}

func NewGameManager() *GameManager {
	return &GameManager{
		settings: &GameSettings{
			Difficulty: "Medium",
			Hints:      "On",
			Control:    "Keyboard",
		},
	}
}

func (gm *GameManager) GetDifficulty() string {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	return gm.settings.Difficulty
}

func (gm *GameManager) GetHints() string {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	return gm.settings.Hints
}
func (gm *GameManager) GetControl() string {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	return gm.settings.Control
}

func (gm *GameManager) SetDifficulty(difficulty string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.settings.Difficulty = difficulty
}

func (gm *GameManager) SetHints(hints string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.settings.Hints = hints
}

func (gm *GameManager) SetControl(control string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.settings.Control = control
}
