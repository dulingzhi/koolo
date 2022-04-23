package run

import (
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/step"
)

type AncientTunnels struct {
	baseRun
}

func (a AncientTunnels) Name() string {
	return "AncientTunnels"
}

func (a AncientTunnels) BuildActions() (actions []step.Runner) {
	// Moving to starting point (Lost City)
	actions = append(actions, a.builder.WayPoint(game.AreaLostCity))

	// Buff
	actions = append(actions, a.char.Buff())

	// Travel to ancient tunnels
	actions = append(actions, step.NewFixedStepsRunner(func(data game.Data) []step.Step {
		return []step.Step{
			step.MoveToLevel(game.AreaAncientTunnels),
		}
	}))

	// Open the chest
	actions = append(actions, step.NewFixedStepsRunner(func(data game.Data) []step.Step {
		for _, o := range data.Objects {
			if o.Name == "SparklyChest" && o.Chest {
				return []step.Step{
					step.MoveTo(o.Position.X, o.Position.Y, true),
					step.InteractObject("SparklyChest", func(data game.Data) bool {
						for _, o := range data.Objects {
							if o.Name == "SparklyChest" && o.Chest {
								return false
							}
						}

						return true
					}),
				}
			}
		}

		return []step.Step{}
	}, step.CanBeSkipped()))

	// Clear Ancient Tunnels
	actions = append(actions, a.char.ClearAncientTunnels())
	return
}
