package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
)

type Service interface {
	FirstOrCreate() *entity.User
}

type service struct {
	repo   Repository
	user   *entity.User
	logger *log.Logger
}

func NewService(user *entity.User) Service {
	return &service{
		repo:   NewRepository(),
		user:   user,
		logger: log.TeeDefault(),
	}
}

func (s *service) FirstOrCreate() *entity.User {
	if result := s.repo.Get(s.user.LineId); result == nil {
		if err := s.repo.Create(s.user); err != nil {
			panic(err.Error())
		}
	} else {
		s.user = result
	}
	return s.user
}
