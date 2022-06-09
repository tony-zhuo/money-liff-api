package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gofrs/uuid"
)

type Service interface {
	GenerateUUIDAndCreateByUser(group *entity.Group, user *entity.User) error
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

func (s *service) GenerateUUIDAndCreateByUser(group *entity.Group, user *entity.User) error {
	u1, err := uuid.NewV1()
	if err != nil {
		return err
	}
	group.UUID = u1.String()
	group.Users = []*entity.User{user}
	group.AdminUserId = user.Id
	if err := s.repo.CreateByUser(group); err != nil {
		return err
	}
	return nil
}
