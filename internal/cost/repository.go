package cost

import (
	"github.com/ZhuoYIZIA/money-liff-api/internal/entity"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCostItemByUser(item *entity.GroupCostItem) error
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

func (r *repository) CreateCostItemByUser(item *entity.GroupCostItem) error {
	result := r.db.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
