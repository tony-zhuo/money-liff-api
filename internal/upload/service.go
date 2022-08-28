package upload

import (
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/storage"
	"mime/multipart"
	"strconv"
	"time"
)

type Service interface {
	UploadImageAndGetPath(file *multipart.FileHeader, folder string) (*string, error)
}

type service struct {
	storage storage.Storage
	logger  *log.Logger
}

func NewService(logger *log.Logger, storage storage.Storage) Service {
	return &service{
		logger:  logger,
		storage: storage,
	}
}

func (s *service) UploadImageAndGetPath(file *multipart.FileHeader, folder string) (*string, error) {
	currentUnixTime := strconv.Itoa(int(time.Now().Unix()))
	fileUrl, err := s.storage.UploadFile(folder, currentUnixTime+"-"+file.Filename, file)
	if err != nil {
		return nil, err
	}

	return fileUrl, nil
}
