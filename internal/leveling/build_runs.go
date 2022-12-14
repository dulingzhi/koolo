package leveling

import (
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/character"
	"github.com/hectorgimenez/koolo/internal/run"
)

func BuildRuns(builder action.Builder, char character.Character) (runs []run.Run) {
	baseRun := run.NewBaseRun(builder, char)

	runs = append(runs, run.Pit{BaseRun: baseRun})

	return
}
