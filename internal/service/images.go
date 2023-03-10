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

func (s *ImagesService) GetAll() ([]uuid.UUID, error) {
	return s.repo.GetAllImagesID()
}

func (s *ImagesService) Get(imageID uuid.UUID, quality int) ([]byte, error) {
	fileName := s.getFileName(imageID.String(), quality)

	img, err := s.repo.Get(fileName)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (s *ImagesService) Create(image image.Image, imageID uuid.UUID, contentType string, quality int) error {
	fileName := s.getFileName(imageID.String(), quality)

	err := s.repo.Create(fileName, contentType, image)
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
