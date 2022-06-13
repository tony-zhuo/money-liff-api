package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
)

type Service interface {
	CreateIfNotFound(user *entity.User) error
	GetUserByLineId(lineId string) *entity.User
}

type service struct {
	repo   Repository
	logger *log.Logger
}

func NewService(repo Repository, logger *log.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) CreateIfNotFound(user *entity.User) error {
	if err := s.repo.FirstOrCreate(user); err != nil {
		return err
	}
	return nil
}

func (s *service) GetUserByLineId(lineId string) *entity.User {
	return s.repo.Get(lineId)
}
