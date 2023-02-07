package run

import (
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/action/step"
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/game/area"
)

const (
	andarielStartingPositionX = 22561
	andarielStartingPositionY = 9553
)

type Andariel struct {
	BaseRun
}

func (a Andariel) Name() string {
	return "Andariel"
}

func (a Andariel) BuildActions() (actions []action.Action) {
	// Moving to starting point (Catacombs Level 2)
	actions = append(actions, a.builder.WayPoint(area.CatacombsLevel2))

	// Buff
	actions = append(actions, a.char.Buff())

	// Travel to boss position
	actions = append(actions, action.BuildStatic(func(data game.Data) []step.Step {
		return []step.Step{
			step.MoveToLevel(area.CatacombsLevel3),
			step.MoveToLevel(area.CatacombsLevel4),
			step.MoveTo(andarielStartingPositionX, andarielStartingPositionY, true),
		}
	}))

	// Kill Andariel
	actions = append(actions, a.char.KillAndariel())
	return
}
