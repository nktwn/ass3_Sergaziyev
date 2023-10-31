package gameengine

type Command interface {
	Execute()
	Undo()
}

type MoveCommand struct {
	player *Player
	x, y   int
}

func NewMoveCommand(player *Player, x, y int) *MoveCommand {
	return &MoveCommand{
		player: player,
		x:      x,
		y:      y,
	}
}

func (cmd *MoveCommand) Execute() {
	cmd.player.MoveTo(cmd.x, cmd.y)
}

func (cmd *MoveCommand) Undo() {
	cmd.player.MoveTo(0, 0)
}

type AttackCommand struct {
	player *Player
}

func NewAttackCommand(player *Player) *AttackCommand {
	return &AttackCommand{
		player: player,
	}
}

func (cmd *AttackCommand) Execute() {
	cmd.player.Attack()
}

func (cmd *AttackCommand) Undo() {
	// отмена атаки
}

type CollectCommand struct {
	player   *Player
	itemName string
}

func NewCollectCommand(player *Player, itemName string) *CollectCommand {
	return &CollectCommand{
		player:   player,
		itemName: itemName,
	}
}

func (cmd *CollectCommand) Execute() {
	cmd.player.CollectItem(cmd.itemName)
}

func (cmd *CollectCommand) Undo() {
	cmd.player.DropItem(cmd.itemName)
}
