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
	DeleteGroupById(id int) error
	GetUserListWithPagination(group *entity.Group, offset, limit int, sort string) (*response.Pagination, error)
	UserJoinGroup(user *entity.User, group *entity.Group) error
	DeleteUserInGroup(group *entity.Group, user *entity.User) error
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

func (s *service) GetListByUserWithPagination(user *entity.User, page, perPage int, sort string) (*response.Pagination, error) {
	groups, err := s.repo.ListByUser(user, page, perPage, sort)
	if err != nil {
		return nil, err
	}
	count := s.repo.GetAllDataCountByUser(user)

	groupResponse := make([]entity.GroupResponse, len(*groups)-1)
	for _, group := range *groups {
		groupResponse = append(groupResponse, entity.GroupResponse{
			UUID:     group.UUID,
			Name:     group.Name,
			ImageUrl: group.ImageUrl,
			IsAdmin:  group.AdminUserId == user.Id,
		})
	}

	pagination := &response.Pagination{
		Page:       page,
		PerPage:    perPage,
		TotalCount: count,
		TotalPage:  int(math.Ceil(float64(count) / float64(perPage))),
		Result:     groupResponse,
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

func (s *service) DeleteGroupById(id int) error {
	if err := s.repo.DeleteGroupById(id); err != nil {
		return err
	}
	return nil
}

func (s *service) GetUserListWithPagination(group *entity.Group, page, perPage int, sort string) (*response.Pagination, error) {
	users, err := s.repo.UserListInGroup(group, page, perPage, sort)
	if err != nil {
		return nil, err
	}

	count := s.repo.GetUserCountInGroup(group)

	pagination := &response.Pagination{
		Page:       page,
		PerPage:    perPage,
		TotalCount: count,
		TotalPage:  int(math.Ceil(float64(count) / float64(perPage))),
		Result:     users,
	}
	return pagination, err
}

func (s *service) UserJoinGroup(user *entity.User, group *entity.Group) error {
	if err := s.repo.AddUserInGroup(group, user); err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteUserInGroup(group *entity.Group, user *entity.User) error {
	if err := s.repo.DeleteUserInGroup(group, user); err != nil {
		return err
	}
	return nil
}
