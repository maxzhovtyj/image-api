package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/maxzhovtyj/image-api/internal/domain"
	"net/http"
)

const maxFileSize = 2 << 20

func (h *Handler) getImage(ctx *gin.Context) {
	// TODO
	panic("implement me")
}

func (h *Handler) addImage(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "form/json")
	err := ctx.Request.ParseMultipartForm(5 << 20)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	files, ok := ctx.Request.MultipartForm.File["file"]
	if len(files) == 0 || !ok {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	fileInfo := files[0]
	fileReader, err := fileInfo.Open()
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if fileInfo.Size > maxFileSize {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	_ = &domain.File{
		Name:   fileInfo.Filename,
		Size:   fileInfo.Size,
		Reader: fileReader,
	}

}
