package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/maxzhovtyj/image-api/internal/service"
	"github.com/maxzhovtyj/image-api/pkg/logger"
)

const (
	apiURL   = "api"
	imageURL = "image"
)

type Handler struct {
	services *service.Service
	logger   logger.Logger
}

func NewHandler(services *service.Service, logger logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	api := router.Group(apiURL)
	{
		api.GET("images-list", h.getImageList)
		api.GET(imageURL, h.getImage)
		api.POST(imageURL, h.addImage)
	}

	return router
}
