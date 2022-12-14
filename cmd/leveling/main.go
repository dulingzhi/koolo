package main

import (
	"context"
	zapLogger "github.com/hectorgimenez/koolo/cmd/koolo/log"
	koolo "github.com/hectorgimenez/koolo/internal"
	"github.com/hectorgimenez/koolo/internal/action"
	"github.com/hectorgimenez/koolo/internal/character"
	"github.com/hectorgimenez/koolo/internal/config"
	"github.com/hectorgimenez/koolo/internal/health"
	"github.com/hectorgimenez/koolo/internal/memory"
	"github.com/hectorgimenez/koolo/internal/run"
	"github.com/hectorgimenez/koolo/internal/town"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %s", err.Error())
	}

	logger, err := zapLogger.NewLogger(config.Config.Debug, config.Config.LogFilePath)
	if err != nil {
		log.Fatalf("Error starting logger: %s", err.Error())
	}
	defer logger.Sync()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	g, ctx := errgroup.WithContext(context.Background())

	process, err := memory.NewProcess(logger)
	if err != nil {
		logger.Fatal("Error finding D2R.exe process", zap.Error(err))
	}

	gr := memory.NewGameReader(process)

	bm := health.NewBeltManager(logger)
	hm := health.NewHealthManager(logger, bm)
	sm := town.NewShopManager(logger, bm)
	char, err := character.BuildCharacter(logger)
	if err != nil {
		logger.Fatal("Error creating character", zap.Error(err))
	}

	ab := action.NewBuilder(logger, sm, bm, gr)
	bot := koolo.NewLevelingBot(logger, hm, ab, gr)
	supervisor := koolo.NewSupervisor(logger, bot, gr)

	g.Go(func() error {
		return supervisor.Start(ctx, run.BuildRuns(ab, char))
	})

	err = g.Wait()
	if err != nil {
		log.Fatalf("Error running Koolo: %s", err.Error())
	}
}
