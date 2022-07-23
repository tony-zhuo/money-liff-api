package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
)

type Service interface {
	RegisterOrFind(user *entity.User) (*entity.User, error)
	GetUserByLineId(lineId string) (*entity.User, error)
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

func (s *service) RegisterOrFind(user *entity.User) (*entity.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return s.repo.FirstOrCreate(user, "line_id = ?", user.LineId)
}

func (s *service) GetUserByLineId(lineId string) (*entity.User, error) {
	return s.repo.Get("line_id = ?", lineId)
}
