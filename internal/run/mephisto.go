package run

import (
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/step"
)

const (
	safeDistanceFromMephistoX = 17568
	safeDistanceFromMephistoY = 8069
)

type Mephisto struct {
	baseRun
}

func (m Mephisto) Name() string {
	return "Mephisto"
}

func (m Mephisto) BuildActions() (actions []step.Runner) {
	// Moving to starting point (Durance of Hate Level 2)
	actions = append(actions, m.builder.WayPoint(game.AreaDuranceOfHateLevel2))

	// Buff
	actions = append(actions, m.char.Buff())

	// Travel to boss position
	actions = append(actions, step.NewFixedStepsRunner(func(data game.Data) []step.Step {
		return []step.Step{
			step.MoveToLevel(game.AreaDuranceOfHateLevel3),
			step.MoveTo(safeDistanceFromMephistoX, safeDistanceFromMephistoY, true),
		}
	}))

	// Kill Mephisto
	actions = append(actions, m.char.KillMephisto())
	return
}
