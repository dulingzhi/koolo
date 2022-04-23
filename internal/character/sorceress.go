package character

import (
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/helper"
	"github.com/hectorgimenez/koolo/internal/hid"
	"github.com/hectorgimenez/koolo/internal/pather"
	"github.com/hectorgimenez/koolo/internal/step"
	"sort"
	"strings"
	"time"
)

const (
	sorceressMaxAttacksLoop = 10
)

type Sorceress struct {
	BaseCharacter
}

func (s Sorceress) Buff() *step.FixedStepsRunner {
	return step.NewFixedStepsRunner(func(data game.Data) (steps []step.Step) {
		steps = append(steps, s.buffCTA()...)
		steps = append(steps, step.SyncStep(func(data game.Data) error {
			if config.Config.Bindings.Sorceress.FrozenArmor != "" {
				hid.PressKey(config.Config.Bindings.Sorceress.FrozenArmor)
				helper.Sleep(100)
				hid.Click(hid.RightButton)
			}

			return nil
		}))

		return
	})
}

func (s Sorceress) KillCountess() *step.FixedStepsRunner {
	return s.killMonster(game.Countess)
}

func (s Sorceress) KillAndariel() *step.FixedStepsRunner {
	return s.killMonster(game.Andariel)
}

func (s Sorceress) KillSummoner() *step.FixedStepsRunner {
	return s.killMonster(game.Summoner)
}

func (s Sorceress) KillPindle() *step.FixedStepsRunner {
	return s.killMonster(game.Pindleskin)
}

func (s Sorceress) KillMephisto() *step.FixedStepsRunner {
	return s.killMonster(game.Mephisto)
}

func (s Sorceress) KillNihlathak() *step.FixedStepsRunner {
	return s.killMonster(game.Nihlathak)
}

func (s Sorceress) ClearAncientTunnels() *step.FixedStepsRunner {
	return nil
}

func (s Sorceress) KillCouncil() *step.RuntimeBuildingRunner {
	return step.NewRuntimeBuildingRunner(func(data game.Data) (step.Step, error) {
		// Exclude monsters that are not council members
		var councilMembers []game.Monster
		var coldImmunes []game.Monster
		for _, m := range data.Monsters {
			if !strings.Contains(strings.ToLower(m.Name), "councilmember") {
				continue
			}
			if m.IsImmune(game.ResistCold) {
				coldImmunes = append(coldImmunes, m)
			} else {
				councilMembers = append(councilMembers, m)
			}
		}

		// Order council members by distance
		sort.Slice(councilMembers, func(i, j int) bool {
			distanceI := pather.DistanceFromPoint(data, councilMembers[i].Position.X, councilMembers[i].Position.Y)
			distanceJ := pather.DistanceFromPoint(data, councilMembers[j].Position.X, councilMembers[j].Position.Y)

			return distanceI < distanceJ
		})

		councilMembers = append(councilMembers, coldImmunes...)

		if len(councilMembers) > 0 {
			m := councilMembers[0]
			return step.NewSecondaryAttack(config.Config.Bindings.Sorceress.Blizzard, game.NPCID(m.Name), 1, time.Second, step.FollowEnemy(3)), nil
			//step.PrimaryAttack(game.NPCID(m.Name), 4, config.Config.Runtime.CastDuration, step.FollowEnemy(maxDistance)),
		}

		return nil, step.ErrNoMoreSteps
	}, step.CanBeSkipped())
}

func (s Sorceress) ClearTrashInArea(distanceFromPlayer int) *step.RuntimeBuildingRunner {
	return step.NewRuntimeBuildingRunner(func(data game.Data) (step.Step, error) {
		return nil, nil
	})
}

func (s Sorceress) killMonster(npc game.NPCID) *step.FixedStepsRunner {
	return step.NewFixedStepsRunner(func(data game.Data) (steps []step.Step) {
		for i := 0; i < sorceressMaxAttacksLoop; i++ {
			steps = append(steps,
				step.NewSecondaryAttack(config.Config.Bindings.Sorceress.Blizzard, npc, 1, time.Second),
				step.PrimaryAttack(npc, 4, config.Config.Runtime.CastDuration),
			)
		}

		return
	}, step.CanBeSkipped())
}
