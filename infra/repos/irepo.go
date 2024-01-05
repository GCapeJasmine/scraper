package repos

import (
	"github.com/cj/scraper/internal/domains/site/repos"
)

type IRepo interface {
	Sites() repos.ISitesRepo
}
