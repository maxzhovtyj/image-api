package service

type Images interface {
}

type Service struct {
	Images
}

func New() *Service {
	return &Service{
		Images: NewImagesService(),
	}
}
