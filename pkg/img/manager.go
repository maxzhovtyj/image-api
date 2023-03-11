package img

import (
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"
)

type ImageManager interface {
	Resize(buf []byte, width, height uint) ([]byte, error)
	Read(path string) ([]byte, error)
	Write(path string, buf []byte) error
}

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Resize(width, height uint, img image.Image) image.Image {
	return resize.Resize(width, height, img, resize.Lanczos3)
}

func (m *Manager) Read(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return img, nil
}

func (m *Manager) Write(path string, image image.Image) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}

	// write new image to file
	err = jpeg.Encode(out, image, nil)
	if err != nil {
		return err
	}

	return nil
}
