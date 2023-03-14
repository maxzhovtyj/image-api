package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/domain"
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

func (r *ImagesRepo) GetAllImagesID() ([]uuid.UUID, error) {
	filesList, err := r.manager.ReadAll(r.dir)
	if err != nil {
		return nil, err
	}

	return filesList, nil
}

func (r *ImagesRepo) Get(fileName string) ([]byte, error) {
	fileRes, err := r.manager.FindFile(r.dir, fileName)
	if err != nil || fileRes == "" {
		return nil, domain.ErrImageNotFound
	}

	readImg, err := r.manager.Read(r.dir + "/" + fileRes)
	if err != nil {
		return nil, err
	}

	return readImg, nil
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
