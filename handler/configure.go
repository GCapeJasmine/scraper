package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/cj/scraper/internal/domains/site/usecases"
)

type Handler struct {
	site *usecases.Site
}

func NewHandler(site *usecases.Site) *Handler {
	return &Handler{
		site: site,
	}
}

func (h *Handler) ConfigAuthRouteAPI(router *gin.RouterGroup) {
	//sites
}
