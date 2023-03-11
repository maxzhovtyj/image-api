package delivery

import "github.com/gin-gonic/gin"

const (
	apiURL   = "api"
	imageURL = "image"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
