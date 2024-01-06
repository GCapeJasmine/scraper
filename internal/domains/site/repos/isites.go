package repos

import (
	"context"
	"github.com/cj/scraper/internal/domain/models"
)

type ISitesRepo interface {
	Create(ctx context.Context, record *models.Site) (*models.Site, error)
	UpdateWithMap(ctx context.Context, record *models.Site, params map[string]interface{}) error
	Upsert(ctx context.Context, record *models.Site) error
	GetSite(ctx context.Context, queries map[string]interface{}) (*models.Site, error)
	GetSiteWithMaximumAccessTime(ctx context.Context) (*models.Site, error)
	GetSiteWithMinimumAccessTime(ctx context.Context) (*models.Site, error)
}
