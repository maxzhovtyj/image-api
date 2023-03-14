package img

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type ImageManager interface {
	FindFile(dir, fileName string) (string, error)
	Resize(width, height uint, img image.Image) image.Image
	ReadAll(dir string) ([]uuid.UUID, error)
	Read(path string) ([]byte, error)
	Write(path string, contentType string, image image.Image) error
}

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

var contentTypeExtension = map[string]string{
	"image/png":  "png",
	"image/jpg":  "jpg",
	"image/jpeg": "jpeg",
}

func (m *Manager) Resize(width, height uint, img image.Image) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
}

func (m *Manager) FindFile(dir, fileName string) (string, error) {
	var fileRes string

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
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
		return "", err
	}

	return fileRes, nil
}

func (m *Manager) ReadAll(dir string) ([]uuid.UUID, error) {
	var res []uuid.UUID

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
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

		contentType := fmt.Sprintf("image/%s", fileExt[1:])

		if contentTypeExtension[contentType] != "" {
			imageNameParts := strings.Split(currFileNameWithoutExt, "_")
			if len(imageNameParts) != 2 {
				return fmt.Errorf("invalid file name")
			}

			parsedImageUUID, err := uuid.Parse(imageNameParts[0])
			if err != nil {
				return err
			}

			res = append(res, parsedImageUUID)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *Manager) Read(path string) ([]byte, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func (m *Manager) Write(path string, contentType string, image image.Image) error {
	out, err := os.Create(fmt.Sprintf("%s.%s", path, contentTypeExtension[contentType]))
	if err != nil {
		return err
	}

	switch contentType {
	case "image/jpeg", "image/jpg":
		err = jpeg.Encode(out, image, nil)
		if err != nil {
			return err
		}
	case "image/png":
		err = png.Encode(out, image)
		if err != nil {
			return err
		}
	}

	return nil
}
