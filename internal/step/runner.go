package step

import (
	"errors"
	"github.com/hectorgimenez/koolo/internal/game"
)

var ErrWillBeRetried = errors.New("error occurred, but it will be retried")
var ErrNoRecover = errors.New("unrecoverable error occurred, game can not continue")
var ErrCanBeSkipped = errors.New("error occurred, but this builder is not critical and game can continue")
var ErrNoMoreSteps = errors.New("builder finished, no more steps remaining")

type Runner interface {
	Next(data game.Data) error
	Skip()
}

type BasicRunner struct {
	retries           int
	canBeSkipped      bool
	resetStepsOnError bool
	markSkipped       bool
}

type Option func(action *BasicRunner)

func CanBeSkipped() Option {
	return func(action *BasicRunner) {
		action.canBeSkipped = true
	}
}

func Resettable() Option {
	return func(action *BasicRunner) {
		action.resetStepsOnError = true
	}
}
