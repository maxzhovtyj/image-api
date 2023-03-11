package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/maxzhovtyj/image-api/internal/service"
)

const (
	apiURL   = "api"
	imageURL = "image"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	api := router.Group(apiURL)
	{
		api.GET(imageURL, h.getImage)
		api.POST(imageURL, h.addImage)
	}

	return router
}
