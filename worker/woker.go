package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/cj/scraper/config"
	"github.com/cj/scraper/internal/domains/site/usecases"
)

type Worker struct {
	cfg         *config.AppConfig
	siteCronJob *usecases.SiteCronJob
}

func NewWorker(cfg *config.AppConfig, siteCronJob *usecases.SiteCronJob) *Worker {
	return &Worker{
		cfg:         cfg,
		siteCronJob: siteCronJob,
	}
}

func (w *Worker) RunJob(ctx context.Context) {
	interval := time.Duration(w.cfg.Worker.Interval) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Job canceled")
			return
		case <-ticker.C:
			fmt.Println("Job running at: ", time.Now())
			w.siteCronJob.UpdateSiteState(ctx)
		}
	}
}
