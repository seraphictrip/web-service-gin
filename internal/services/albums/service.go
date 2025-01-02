package albums

import (
	"errors"
	"example/web-service-gin/internal/models"
	"example/web-service-gin/internal/repositories"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrInternalService = errors.New("internal service erro")
)

type AlbumsService interface {
	Get() ([]models.Album, error)
	Create(models.Album) error
	GetById(id string) (models.Album, error)
}

type svc struct {
	repo repositories.AlbumsRepository
}

func NewAlbumsService(repo repositories.AlbumsRepository) *svc {
	return &svc{repo}
}
func (s *svc) Get() ([]models.Album, error) {
	// TODO: if repo can throw eror we should trap and convert to business appropriate error
	return s.repo.Get()
}

func (s *svc) Create(newAlbum models.Album) error {
	// TODO: if repo can throw eror we should trap and convert to business appropriate error
	err := s.repo.Create(newAlbum)
	return err
}

func (s *svc) GetById(id string) (album models.Album, err error) {
	album, err = s.repo.GetById(id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			err = ErrNotFound
		} else {
			err = ErrInternalService
		}
	}
	return album, err
}
