package main

import (
	"flag"

	"go.uber.org/zap"

	"github.com/cj/scraper/config"
	"github.com/cj/scraper/server"
	"github.com/cj/scraper/setting"
)

func main() {
	var configFile, port string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.StringVar(&port, "port", "", "Specify port")
	flag.Parse()

	defer setting.WaitOSSignal()

	//load config
	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("load config fail with err: %v", err)
		panic(err)
	}

	// migrate db
	go setting.MigrateDatabase(cfg.DB)

	//start new server
	s := server.NewServer(cfg)
	s.Init()

	if err := s.ListenHTTP(); err != nil {
		zap.S().Errorf("start server fail with err: %v", err)
		panic(err)
	}
}
