package settings

type Settings interface {
	GetDifficulty() string
	GetHints() string
	GetControl() string
	SetDifficulty(difficulty string)
	SetHints(hints string)
	SetControl(control string)
}
