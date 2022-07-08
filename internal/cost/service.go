package cost

import "github.com/ZhuoYIZIA/money-liff-api/pkg/log"

type Service interface {
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