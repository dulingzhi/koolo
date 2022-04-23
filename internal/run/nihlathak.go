package run

import (
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/step"
)

type Nihlathak struct {
	baseRun
}

func (a Nihlathak) Name() string {
	return "Nihlathak"
}

func (a Nihlathak) BuildActions() (actions []step.Runner) {
	// Moving to starting point (Halls of Pain)
	actions = append(actions, a.builder.WayPoint(game.AreaHallsOfPain))

	// Buff
	actions = append(actions, a.char.Buff())

	// Travel to boss position
	actions = append(actions, step.NewFixedStepsRunner(func(data game.Data) []step.Step {
		return []step.Step{
			step.MoveToLevel(game.AreaHallsOfVaught),
		}
	}))

	// Move to Nilhatak
	actions = append(actions, step.NewFixedStepsRunner(func(data game.Data) []step.Step {
		for _, n := range data.PointsOfInterest {
			if n.Name == a.Name() {
				return []step.Step{step.MoveTo(n.Position.X, n.Position.Y, true)}
			}
		}

		return []step.Step{}
	}))

	// Kill Nihlathak
	actions = append(actions, a.char.KillNihlathak())
	return
}
