package server

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/cj/scraper/config"
	"github.com/cj/scraper/infra"
	"github.com/cj/scraper/infra/repos"
)

// Server ...
type Server struct {
	httpServer *http.Server
	router     *gin.Engine
	cfg        *config.AppConfig
}

func NewServer(cfg *config.AppConfig) *Server {
	router := gin.New()
	return &Server{
		router: router,
		cfg:    cfg,
	}
}

func (s *Server) Init() {
	db, err := infra.InitMySQL(s.cfg.DB)
	if err != nil {
		zap.S().Errorf("Init db error: %v", err)
		panic(err)
	}
	repo := repos.NewSQLRepo(db, s.cfg.DB)
	domains := s.initDomains(repo)
	s.initCORS()
	s.initRouter(domains, repo)
}

// ListenHTTP ...
func (s *Server) ListenHTTP() error {
	listen, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		zap.S().Errorf("err %v", err)
		panic(err)
	}
	address := fmt.Sprintf(":%s", os.Getenv("PORT"))

	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	zap.S().Infof("starting http server at port %v ...", os.Getenv("PORT"))

	return s.httpServer.Serve(listen)
}
