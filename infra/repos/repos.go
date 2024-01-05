package repos

import (
	"gorm.io/gorm"

	"github.com/cj/scraper/config"
	"github.com/cj/scraper/internal/domains/site/repos"
)

// Repo ...
type Repo struct {
	db  *gorm.DB
	cfg *config.MySQLConfig
}

// NewSQLRepo ...
func NewSQLRepo(db *gorm.DB, cfg *config.MySQLConfig) IRepo {
	return &Repo{
		db:  db,
		cfg: cfg,
	}
}

func (r *Repo) Sites() repos.ISitesRepo {
	return NewSitesSQLRepo(r.db)
}
