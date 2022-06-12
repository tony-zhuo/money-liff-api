package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/database"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
	ListByUser(user *entity.User, offset, limit int, sort string) (*[]entity.Group, error)
	CreateByUser(group *entity.Group, user *entity.User) error
	GetAllDataCountByUser(user *entity.User) int
}

type repository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRepository() Repository {
	return &repository{
		db:     database.Connection(),
		logger: log.TeeDefault(),
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

func (r *repository) CreateByUser(group *entity.Group, user *entity.User) error {
	group.Users = []*entity.User{user}
	group.AdminUserId = user.Id
	if result := r.db.Create(group); result.Error != nil {
		return result.Error
	}
	return nil
}
