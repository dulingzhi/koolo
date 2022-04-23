package run

import (
	"github.com/hectorgimenez/koolo/internal/character"
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/step"
	"github.com/hectorgimenez/koolo/internal/step/builder"
	"strings"
)

type Run interface {
	Name() string
	BuildActions() []step.Runner
}

type baseRun struct {
	builder builder.Builder
	char    character.Character
}

func BuildRuns(builder builder.Builder, char character.Character) (runs []Run) {
	baseRun := baseRun{
		builder: builder,
		char:    char,
	}

	for _, run := range config.Config.Game.Runs {
		run = strings.ToLower(run)
		switch run {
		case "countess":
			runs = append(runs, Countess{baseRun})
		case "andariel":
			runs = append(runs, Andariel{baseRun})
		case "summoner":
			runs = append(runs, Summoner{baseRun})
		case "mephisto":
			runs = append(runs, Mephisto{baseRun})
		case "council":
			runs = append(runs, Council{baseRun})
		case "pindleskin":
			runs = append(runs, Pindleskin{baseRun})
		case "nihlathak":
			runs = append(runs, Nihlathak{baseRun})
		case "ancient_tunnels":
			runs = append(runs, AncientTunnels{baseRun})
		}
	}

	return
}
