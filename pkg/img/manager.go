package img

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

type ImageManager interface {
	Resize(width, height uint, img image.Image) image.Image
	Read(path string) ([]byte, error)
	Write(path string, contentType string, image image.Image) error
}

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Resize(width, height uint, img image.Image) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
}

func (m *Manager) Read(path string) ([]byte, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

var contentTypeExtension = map[string]string{
	"image/png":  "png",
	"image/jpg":  "jpg",
	"image/jpeg": "jpeg",
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
