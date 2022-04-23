package character

import (
	"fmt"
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/game"
	"github.com/hectorgimenez/koolo/internal/helper"
	"github.com/hectorgimenez/koolo/internal/hid"
	"github.com/hectorgimenez/koolo/internal/step"
	"strings"
)

type Character interface {
	Buff() *step.FixedStepsRunner
	KillCountess() *step.FixedStepsRunner
	KillAndariel() *step.FixedStepsRunner
	KillSummoner() *step.FixedStepsRunner
	KillMephisto() *step.FixedStepsRunner
	KillPindle() *step.FixedStepsRunner
	KillNihlathak() *step.FixedStepsRunner
	KillCouncil() *step.RuntimeBuildingRunner
	ClearAncientTunnels() *step.FixedStepsRunner
	ClearTrashInArea(distanceFromPlayer int) *step.RuntimeBuildingRunner
}

func BuildCharacter() (Character, error) {
	bc := BaseCharacter{}
	switch strings.ToLower(config.Config.Character.Class) {
	case "sorceress":
		return Sorceress{BaseCharacter: bc}, nil
	case "hammerdin":
		return Hammerdin{BaseCharacter: bc}, nil
	}

	return nil, fmt.Errorf("class %s not implemented", config.Config.Character.Class)
}

type BaseCharacter struct {
}

func (bc BaseCharacter) buffCTA() (steps []step.Step) {
	if config.Config.Character.UseCTA {
		steps = append(steps,
			step.SwapWeapon(),
			step.SyncStep(func(data game.Data) error {
				helper.Sleep(1000)
				hid.PressKey(config.Config.Bindings.CTABattleCommand)
				helper.Sleep(100)
				hid.Click(hid.RightButton)
				helper.Sleep(500)
				hid.PressKey(config.Config.Bindings.CTABattleOrders)
				helper.Sleep(100)
				hid.Click(hid.RightButton)
				helper.Sleep(1000)

				return nil
			}),
			step.SwapWeapon(),
		)
	}

	return steps
}
