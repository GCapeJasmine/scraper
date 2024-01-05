package repos

import (
	"context"
	"github.com/cj/scraper/internal/domain/models"
)

type ISitesRepo interface {
	Create(ctx context.Context, record *models.Site) (*models.Site, error)
	UpdateWithMap(ctx context.Context, record *models.Site, params map[string]interface{}) error
}
