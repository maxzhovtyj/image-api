package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/pkg/img"
	"image"
)

type ImagesRepo struct {
	dir     string
	manager img.ImageManager
}

func NewImagesRepo(imageDirPath string, manager img.ImageManager) *ImagesRepo {
	return &ImagesRepo{
		dir:     imageDirPath,
		manager: manager,
	}
}

func (r *ImagesRepo) Get(imageID uuid.UUID, quality int) (image.Image, error) {
	return nil, nil
}

func (r *ImagesRepo) Create(name string, contentType string, image image.Image) error {
	filePath := fmt.Sprintf("%s/%s", r.dir, name)

	err := r.manager.Write(filePath, contentType, image)
	if err != nil {
		return err
	}

	return nil
}

func (r *ImagesRepo) Resize(width, height uint, img image.Image) image.Image {
	return r.manager.Resize(width, height, img)
}
