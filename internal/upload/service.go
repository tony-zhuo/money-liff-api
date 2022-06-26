package upload

import (
	"bytes"
	"github.com/ZhuoYIZIA/money-liff-api/pkg/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

type Service interface {
	UploadImageAndGetPath(file *multipart.FileHeader, folder string) (*string, error)
}

type service struct {
	logger *log.Logger
}

func NewService(logger *log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) UploadImageAndGetPath(file *multipart.FileHeader, folder string) (*string, error) {
	region := os.Getenv("AWS_S3_REGION")
	bucket := os.Getenv("AWS_S3_BUCKET")
	s3Key := os.Getenv("AWS_S3_KEY")
	s3Secret := os.Getenv("AWS_S3_SECRET")
	s3Endpoint := os.Getenv("AWS_S3_ENDPOINT")
	credential := credentials.NewStaticCredentials(s3Key, s3Secret, "")
	currentUnixTime := strconv.Itoa(int(time.Now().Unix()))

	s.logger.Info("UploadImageAndGetPath currentTimeString", log.String("time", currentUnixTime))
	fileEndpoint := folder + "/" + currentUnixTime + "-" + file.Filename
	s.logger.Info("UploadImageAndGetPath fileEndpoint", log.String("endpoint", fileEndpoint))

	size := file.Size

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credential,
	})
	if err != nil {
		s.logger.Error("UploadImageAndGetPath new session error", log.String("err", err.Error()))
		return nil, err
	}

	f, err := file.Open()
	if err != nil {
		s.logger.Error("UploadImageAndGetPath file open error", log.String("err", err.Error()))
		return nil, err
	}

	buffer := make([]byte, size)
	_, err = f.Read(buffer)
	if err != nil {
		s.logger.Error("UploadImageAndGetPath file read error", log.String("err", err.Error()))
		return nil, err
	}

	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileEndpoint),
		Body:   bytes.NewReader(buffer),
	})
	if err != nil {
		s.logger.Error("UploadImageAndGetPath put object error", log.String("err", err.Error()))
		return nil, err
	}

	fileUrl := s3Endpoint + "/" + fileEndpoint
	return &fileUrl, nil
}
