package delivery

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/domain"
	"io"
	"net/http"
	"strconv"
)

const (
	maxFileSize         = 2 << 20
	defaultImageQuality = 100
)

func (h *Handler) getImage(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/octet-stream")

	imageIDParam := ctx.Query("id")
	imageQualityParam := ctx.Query("quality")

	imageQuality, err := strconv.Atoi(imageQualityParam)
	if err != nil {
		if imageQualityParam == "" {
			imageQuality = defaultImageQuality
		} else {
			newErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid image quality").Error())
			return
		}
	}

	imageUUID, err := uuid.Parse(imageIDParam)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("invalid image uuid").Error())
		return
	}

	imageBytes, err := h.services.Images.Get(imageUUID, imageQuality)
	if err != nil {
		if errors.Is(err, domain.ErrImageNotFound) {
			newErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}

		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = ctx.Writer.Write(imageBytes)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getImageList(ctx *gin.Context) {
	all, err := h.services.Images.GetAll()
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, all)
}

func (h *Handler) addImage(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(maxFileSize)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	files, ok := ctx.Request.MultipartForm.File["file"]
	if len(files) == 0 || !ok {
		newErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("file not valid").Error())
		return
	}

	fileInfo := files[0]
	fileReader, err := fileInfo.Open()
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("failed to read file %v", err).Error())
		return
	}

	if fileInfo.Size > maxFileSize {
		newErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("file size exceeded").Error())
		return
	}

	imageBuf := bytes.NewBuffer(nil)
	_, err = io.Copy(imageBuf, fileReader)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, fmt.Errorf("failed to convert file to bytes buffer").Error())
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

	ctx.String(http.StatusCreated, "image successfully published")
}
