package step

import (
	"fmt"
	"github.com/hectorgimenez/koolo/internal/game"
)

const maxRetries = 5

type FixedStepsRunner struct {
	Steps           []Step
	builder         func(data game.Data) []Step
	builderExecuted bool
	*BasicRunner
}

func NewFixedStepsRunner(builder func(data game.Data) []Step, opts ...Option) *FixedStepsRunner {
	r := &FixedStepsRunner{
		builder:     builder,
		BasicRunner: &BasicRunner{},
	}

	for _, opt := range opts {
		opt(r.BasicRunner)
	}

	return r
}

func (r *FixedStepsRunner) resetSteps() {
	if !r.resetStepsOnError {
		return
	}

	for _, s := range r.Steps {
		s.Reset()
	}
}

func (r *BasicRunner) Skip() {
	if r.retries >= maxRetries && r.canBeSkipped {
		r.markSkipped = true
	}
}

func (r *FixedStepsRunner) Next(data game.Data) error {
	if r.markSkipped {
		return ErrNoMoreSteps
	}

	if r.retries >= maxRetries {
		if r.canBeSkipped {
			return fmt.Errorf("%w: attempt limit reached", ErrCanBeSkipped)
		}
		return fmt.Errorf("%w: attempt limit reached", ErrNoRecover)
	}

	if r.builder != nil && !r.builderExecuted {
		r.Steps = r.builder(data)
		r.builderExecuted = true
	}

	for _, s := range r.Steps {
		if s.Status(data) != StatusCompleted {
			err := s.Run(data)
			if err != nil {
				r.retries++
				r.resetSteps()
			}

			return nil
		}
	}

	return ErrNoMoreSteps
}
