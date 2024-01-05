package setting

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/cj/scraper/config"
	"github.com/cj/scraper/infra"
)

const (
	migrationFile = "file://./migrations/sql"
)

// MigrateDatabase ...
func MigrateDatabase(cfg *config.MySQLConfig) {
	infra.CreateDBAndMigrate(cfg, migrationFile)
}

// WaitOSSignal ...
func WaitOSSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	s := <-c
	zap.S().Infof("Receive os.Signal: %s", s.String())
}
