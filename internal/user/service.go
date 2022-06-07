package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
)

type Service interface {
	CreateIfNotFound(user *entity.User) error
}

type service struct {
	repo   Repository
	logger *log.Logger
}

func NewService() Service {
	return &service{
		repo:   NewRepository(),
		logger: log.TeeDefault(),
	}
}

func (s *service) CreateIfNotFound(user *entity.User) error {
	if err := s.repo.FirstOrCreate(user); err != nil {
		return err
	}
	return nil
}
