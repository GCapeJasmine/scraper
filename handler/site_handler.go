package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cj/scraper/handler/models"
	modelsUsecase "github.com/cj/scraper/internal/domain/models"
)

func (h *Handler) getSite() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := &models.GetSiteRequest{}
		if err := c.ShouldBind(request); err != nil {
			log.Printf("parse request with err = %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := h.site.GetSite(c, &modelsUsecase.GetSiteInput{
			Name:                request.Name,
			IsMaximumAccessTime: request.IsMaximumAccessTime,
			IsMinimumAccessTime: request.IsMinimumAccessTime,
		})
		if err != nil {
			log.Printf("GetSiteUsecase fail with err = %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
