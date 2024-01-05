package server

import (
	"github.com/gin-contrib/cors"

	hdl "github.com/cj/scraper/handler"
	"github.com/cj/scraper/infra/repos"
	"github.com/cj/scraper/internal/domains/site/usecases"
)

type domains struct {
	site *usecases.Site
}

func (s *Server) initCORS() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Access-Token",
		"X-Google-Access-Token",
	}
	s.router.Use(cors.New(corsConfig))
}

func (s *Server) initDomains(repo repos.IRepo) *domains {
	site := usecases.NewSite(repo.Sites())
	return &domains{
		site: site,
	}
}

func (s *Server) initRouter(domains *domains, repo repos.IRepo) {
	// init handler
	handler := hdl.NewHandler(domains.site)

	// Auth API
	routerAuth := s.router.Group("v1")
	handler.ConfigAuthRouteAPI(routerAuth)
}
