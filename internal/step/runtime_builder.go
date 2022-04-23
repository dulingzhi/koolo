package step

import "github.com/hectorgimenez/koolo/internal/game"

type RuntimeBuildingRunner struct {
	stepBuilderFn func(game.Data) (Step, error)
	runningStep   Step
	*BasicRunner
}

func NewRuntimeBuildingRunner(stepBuilderFn func(game.Data) (Step, error), opts ...Option) *RuntimeBuildingRunner {
	r := &RuntimeBuildingRunner{
		stepBuilderFn: stepBuilderFn,
		BasicRunner:   &BasicRunner{},
	}

	for _, opt := range opts {
		opt(r.BasicRunner)
	}

	return r
}

func (rn *RuntimeBuildingRunner) Next(data game.Data) error {
	if rn.runningStep == nil || rn.runningStep.Status(data) == StatusCompleted {
		step, err := rn.stepBuilderFn(data)
		if err != nil {
			return err
		}
		rn.runningStep = step
	}

	if rn.runningStep.Status(data) != StatusCompleted {
		err := rn.runningStep.Run(data)
		if err != nil {
			rn.retries++
		}

		return nil
	}

	return ErrNoMoreSteps
}
