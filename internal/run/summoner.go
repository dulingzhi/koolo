package run

import (
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/step"
)

type Summoner struct {
	baseRun
}

func (s Summoner) Name() string {
	return "Summoner"
}

func (s Summoner) BuildActions() (actions []step.Runner) {
	// Moving to starting point (Arcane Sanctuary)
	actions = append(actions, s.builder.WayPoint(game.AreaArcaneSanctuary))

	// Buff
	actions = append(actions, s.char.Buff())

	// Travel to boss position
	actions = append(actions, step.NewFixedStepsRunner(func(data game.Data) []step.Step {
		npc, found := data.NPCs.FindOne(game.Summoner)
		if !found {
			return nil
		}

		return []step.Step{
			step.MoveTo(npc.Positions[0].X, npc.Positions[0].Y, true),
		}
	}, step.CanBeSkipped()))

	// Kill Summoner
	actions = append(actions, s.char.KillSummoner())
	return
}
