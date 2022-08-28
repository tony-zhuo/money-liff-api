package storage

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
	"os"
)

type Storage interface {
	UploadFile(path string, filename string, file *multipart.FileHeader) (fileUrl *string, err error)
}

type storage struct {
	region     string
	bucket     string
	s3Endpoint string
	credential *credentials.Credentials
}

func NewStorage() Storage {
	region := os.Getenv("AWS_S3_REGION")
	bucket := os.Getenv("AWS_S3_BUCKET")
	s3Key := os.Getenv("AWS_S3_KEY")
	s3Secret := os.Getenv("AWS_S3_SECRET")
	s3Endpoint := os.Getenv("AWS_S3_ENDPOINT")

	return &storage{
		region:     region,
		bucket:     bucket,
		s3Endpoint: s3Endpoint,
		credential: credentials.NewStaticCredentials(s3Key, s3Secret, ""),
	}
}

func (s *storage) UploadFile(path string, filename string, file *multipart.FileHeader) (fileUrl *string, err error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s.region),
		Credentials: s.credential,
	})
	if err != nil {
		return nil, err
	}

	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buffer := make([]byte, file.Size)
	_, err = f.Read(buffer)
	if err != nil {
		return nil, err
	}

	fileEndpoint := path + "/" + filename

	_, err = s3.New(sess).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileEndpoint),
		Body:   bytes.NewReader(buffer),
	})

	url := s.s3Endpoint + "/" + fileEndpoint

	return &url, nil
}
