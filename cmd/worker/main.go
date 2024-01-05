package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cj/scraper/config"
	"github.com/cj/scraper/infra"
	"github.com/cj/scraper/infra/repos"
	"github.com/cj/scraper/internal/domains/site/usecases"
	"github.com/cj/scraper/setting"
	"github.com/cj/scraper/worker"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.Parse()

	defer setting.WaitOSSignal()

	//load config
	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("load config fail with err: %v", err)
		panic(err)
	}

	db, err := infra.InitMySQL(cfg.DB)
	if err != nil {
		zap.S().Errorf("Init db error: %v", err)
		panic(err)
	}

	//init repo
	repo := repos.NewSQLRepo(db, cfg.DB)
	if err != nil {
		zap.S().Errorf("Init repo error: %v", err)
		panic(err)
	}
	siteConJob := usecases.NewSiteCronJob(repo.Sites())
	w := worker.NewWorker(cfg, siteConJob)
	ctx, cancel := context.WithCancel(context.Background())
	w.RunJob(ctx)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	fmt.Println("Received interrupt signal, canceling job...")
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Exiting...")
}
