package middleware

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	Client *s3.S3
	Bucket string
}

func NewS3Service(client *s3.S3, bucket string) *S3Service {
	return &S3Service{Client: client, Bucket: bucket}
}

func (s *S3Service) GenerateUploadURL(filename string) (string, error) {
	req, _ := s.Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
	})

	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *S3Service) GenerateDownloadURL(filename string) (string, error) {
	req, _ := s.Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
	})

	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil
}
