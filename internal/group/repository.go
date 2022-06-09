package group

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/database"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
	CreateByUser(group *entity.Group) error
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

func (r *repository) CreateByUser(group *entity.Group) error {
	if result := r.db.Create(group); result.Error != nil {
		return result.Error
	}
	return nil
}
