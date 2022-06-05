package user

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
)

type Service interface {
	FirstOrCreate() *entity.User
}

type service struct {
	repo Repository
	user *entity.User
}

func NewService(user *entity.User) Service {
	return &service{
		repo: NewRepository(),
		user: user,
	}
}

func (s *service) FirstOrCreate() *entity.User {
	result := s.repo.Get(s.user.LineId)
	if result == nil {
		err := s.repo.Create(s.user)
		if err != nil {
			//TODO: error handle
		}
	} else {
		s.user = result
	}
	return s.user
}
