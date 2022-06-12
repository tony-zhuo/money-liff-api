package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/internal/unity/response"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/gofrs/uuid"
	"math"
)

type Service interface {
	GetListByUserWithPagination(user *entity.User, offset, limit int, sort string) (*response.Pagination, error)
	GenerateUUIDAndCreateByUser(group *entity.Group, user *entity.User) error
	GetGroupByUUID(uuid string) *entity.Group
	CheckUserIsAdmin(group *entity.Group, user *entity.User) bool
	UpdateGroupById(group *entity.Group, id int) error
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

func (s *service) GetListByUserWithPagination(user *entity.User, page, perPage int, sort string) (*response.Pagination, error) {
	groups, err := s.repo.ListByUser(user, page, perPage, sort)
	if err != nil {
		return nil, err
	}
	count := s.repo.GetAllDataCountByUser(user)

	pagination := &response.Pagination{
		Page:       page,
		PerPage:    perPage,
		TotalCount: count,
		TotalPage:  int(math.Ceil(float64(count) / float64(perPage))),
		Result:     groups,
	}
	return pagination, nil
}

func (s *service) GenerateUUIDAndCreateByUser(group *entity.Group, user *entity.User) error {
	u1, err := uuid.NewV1()
	if err != nil {
		return err
	}
	group.UUID = u1.String()
	if err := s.repo.CreateByUser(group, user); err != nil {
		return err
	}
	return nil
}

func (s *service) GetGroupByUUID(uuid string) *entity.Group {
	return s.repo.GetByUUID(uuid)
}

func (s *service) CheckUserIsAdmin(group *entity.Group, user *entity.User) bool {
	return group.AdminUserId == user.Id
}

func (s *service) UpdateGroupById(group *entity.Group, id int) error {
	if err := s.repo.UpdateGroupById(group, id); err != nil {
		return err
	}
	return nil
}
