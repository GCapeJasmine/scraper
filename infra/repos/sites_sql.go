package repos

import (
	"context"

	"gorm.io/gorm"

	"github.com/cj/scraper/internal/domain/models"
)

type sitesSQLRepo struct {
	db *gorm.DB
}

func NewSitesSQLRepo(db *gorm.DB) *sitesSQLRepo {
	return &sitesSQLRepo{
		db: db,
	}
}

func (s *sitesSQLRepo) Create(ctx context.Context, record *models.Site) (*models.Site, error) {
	err := s.db.Create(record).Error
	return record, err
}

func (s *sitesSQLRepo) UpdateWithMap(ctx context.Context, record *models.Site, params map[string]interface{}) error {
	return nil
}
