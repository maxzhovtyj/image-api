package repository

import (
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/pkg/img"
	"image"
	"log"
	"os"
)

type Images interface {
	Get(imageID uuid.UUID, quality int) (image.Image, error)
	Create(name string, image image.Image) error
	Resize(width, height uint, img image.Image) image.Image
}

type Repository struct {
	Images
}

func New(imgDirPath string, manager img.ImageManager) *Repository {
	_, err := os.Stat(imgDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(imgDirPath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return &Repository{
		Images: NewImagesRepo(imgDirPath, manager),
	}
}
