package run

import (
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/action/step"
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/game/area"
)

const ()

type Pit struct {
	BaseRun
}

func (a Pit) Name() string {
	return "Pit"
}

func (a Pit) BuildActions() (actions []action.Action) {
	// Moving to starting point (RogueEncampment)
	actions = append(actions, a.builder.WayPoint(area.RogueEncampment))

	// Buff
	actions = append(actions, a.char.Buff())

	// Travel to boss position
	actions = append(actions, action.BuildStatic(func(data game.Data) []step.Step {
		return []step.Step{
			step.MoveToLevel(area.DenOfEvil),
		}
	}))

	// Kill Andariel
	actions = append(actions, a.char.Ki())
	return
}
