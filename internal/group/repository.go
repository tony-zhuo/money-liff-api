package group

import (
	"errors"
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
	ListByUser(user *entity.User, offset, limit int, sort string) (*[]entity.Group, error)
	CreateByUser(group *entity.Group, user *entity.User) error
	GetByUUID(uuid string) *entity.Group
	GetAllDataCountByUser(user *entity.User) int
	UpdateGroupById(group *entity.Group, id int) error
	DeleteGroupById(id int) error
	GetUserListInGroup(group *entity.Group) (*[]entity.User, error)
	AddUserInGroup(group *entity.Group, user *entity.User) error
}

type repository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRepository(db *gorm.DB, logger *log.Logger) Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (r *repository) ListByUser(user *entity.User, page, perPage int, sort string) (*[]entity.Group, error) {
	var groups *[]entity.Group

	err := r.db.
		Model(user).
		Order(sort).
		Offset((page - 1) * perPage).
		Limit(perPage).
		Association("Groups").
		Find(&groups)

	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *repository) GetAllDataCountByUser(user *entity.User) int {
	count := r.db.
		Model(user).
		Association("Groups").
		Count()
	return int(count)
}

func (r *repository) GetByUUID(uuid string) *entity.Group {
	group := entity.Group{}
	result := r.db.Where("uuid = ?", uuid).First(&group)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &group
}

func (r *repository) CreateByUser(group *entity.Group, user *entity.User) error {
	group.Users = []*entity.User{user}
	group.AdminUserId = user.Id
	if result := r.db.Create(group); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) UpdateGroupById(group *entity.Group, id int) error {
	if result := r.db.Model(&entity.Group{Id: id}).Updates(group); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) DeleteGroupById(id int) error {
	if result := r.db.Delete(&entity.Group{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) GetUserListInGroup(group *entity.Group) (*[]entity.User, error) {
	var users *[]entity.User
	err := r.db.
		Model(group).
		Association("users").
		Find(users)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) AddUserInGroup(group *entity.Group, user *entity.User) error {
	err := r.db.Model(user).Association("Groups").Append([]*entity.Group{group})
	if err != nil {
		return err
	}
	return nil
}
