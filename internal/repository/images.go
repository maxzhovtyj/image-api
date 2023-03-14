package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/maxzhovtyj/image-api/internal/domain"
	"github.com/maxzhovtyj/image-api/pkg/img"
	"image"
	"io/fs"
	"path/filepath"
	"strings"
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
	var fileRes string

	err := filepath.Walk(r.dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		currFileName := info.Name()

		if info.IsDir() {
			return nil
		}

		fileBase := filepath.Base(currFileName)
		fileExt := filepath.Ext(currFileName)

		currFileNameWithoutExt := strings.TrimSuffix(fileBase, fileExt)

		if currFileNameWithoutExt == fileName {
			fileRes = currFileName
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	if fileRes == "" {
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
