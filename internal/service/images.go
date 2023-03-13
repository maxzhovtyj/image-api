package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/repository"
	"image"
)

type ImagesService struct {
	repo repository.Images
}

func NewImagesService(images repository.Images) *ImagesService {
	return &ImagesService{repo: images}
}

func (s *ImagesService) getFileName(imageID string, quality int) string {
	return fmt.Sprintf("%s_%d", imageID, quality)
}

func (s *ImagesService) Get(imageID uuid.UUID, quality int) (image.Image, error) {
	return nil, nil
}

func (s *ImagesService) Create(image image.Image, contentType string, quality int) error {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	fileName := s.getFileName(newUUID.String(), quality)

	err = s.repo.Create(fileName, contentType, image)
	if err != nil {
		return err
	}

	return nil
}

func (s *ImagesService) Resize(img image.Image, width, height int) (image.Image, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("invalid image sizes")
	}

	resizedImage := s.repo.Resize(
		uint(width),
		uint(height),
		img,
	)

	return resizedImage, nil
}
