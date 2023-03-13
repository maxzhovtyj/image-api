package delivery

import (
	"bytes"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/domain"
	"io"
	"net/http"
	"strconv"
)

const maxFileSize = 2 << 20

func (h *Handler) getImage(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/octet-stream")
	imageID := ctx.Query("id")
	imageQuality := ctx.Query("quality")

	imageQualityInt, err := strconv.Atoi(imageQuality)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	imageUUID, err := uuid.Parse(imageID)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	imageBytes, err := h.services.Images.Get(imageUUID, imageQualityInt)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Writer.Write(imageBytes)
}

func (h *Handler) addImage(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "form/json")
	err := ctx.Request.ParseMultipartForm(maxFileSize)
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

	imageBuf := bytes.NewBuffer(nil)
	_, err = io.Copy(imageBuf, fileReader)
	if err != nil {
		return
	}

	err = h.services.Publisher.Publish(context.Background(), imageBuf.Bytes())
	if err != nil {
		if errors.Is(err, domain.ErrInvalidContentType) {
			newErrorResponse(ctx, http.StatusUnsupportedMediaType, err.Error())
			return
		}

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusOK, "image successfully published")
}
