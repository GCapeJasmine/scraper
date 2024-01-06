package repos

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

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

func (s *sitesSQLRepo) Upsert(ctx context.Context, record *models.Site) error {
	return s.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"access_time", "state"}),
	}).Create(record).Error
}

func (s *sitesSQLRepo) GetSite(ctx context.Context, queries map[string]interface{}) (*models.Site, error) {
	site := &models.Site{}
	query := s.db
	if len(queries) > 0 {
		query = query.Where(queries)
	}
	err := query.First(&site).Error
	return site, err
}

func (s *sitesSQLRepo) GetSiteWithMaximumAccessTime(ctx context.Context) (*models.Site, error) {
	res := &models.Site{}
	err := s.db.Where("access_time = (?)", s.db.Table("sites").Select("MAX(access_time)")).First(&res).Error
	return res, err
}

func (s *sitesSQLRepo) GetSiteWithMinimumAccessTime(ctx context.Context) (*models.Site, error) {
	res := &models.Site{}
	err := s.db.Where("access_time = (?)", s.db.Table("sites").Select("MIN(access_time)")).First(&res).Error
	return res, err
}
