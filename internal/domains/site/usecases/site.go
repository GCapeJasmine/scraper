package usecases

import (
	"github.com/cj/scraper/internal/domains/site/repos"
)

type Site struct {
	sitesRepo repos.ISitesRepo
}

func NewSite(sitesRepo repos.ISitesRepo) *Site {
	return &Site{
		sitesRepo: sitesRepo,
	}
}
