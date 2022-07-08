package cost

import (
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"gorm.io/gorm"
)

type Repository interface {
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
