package usecases

import (
	"context"
	"github.com/cj/scraper/internal/domain/models"
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

func (s *Site) GetSite(ctx context.Context, input *models.GetSiteInput) (*models.Site, error) {
	name, isMaximumAccessTime, isMinimumAccessTime := input.Name, input.IsMaximumAccessTime, input.IsMinimumAccessTime
	queries := make(map[string]interface{})
	if len(name) > 0 {
		queries["name"] = name
		res, err := s.sitesRepo.GetSite(ctx, queries)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else if isMaximumAccessTime {
		res, err := s.sitesRepo.GetSiteWithMaximumAccessTime(ctx)
		if err != nil {
			return nil, err
		}
		return res, err
	} else if isMinimumAccessTime {
		res, err := s.sitesRepo.GetSiteWithMinimumAccessTime(ctx)
		if err != nil {
			return nil, err
		}
		return res, err
	}
	return nil, nil
}
