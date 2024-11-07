package middleware

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service struct {
	Client *s3.Client
	Bucket string
}

func NewS3Service(client *s3.Client, bucket string) *S3Service {
	return &S3Service{Client: client, Bucket: bucket}
}

func (s *S3Service) GenerateUploadURL(ctx context.Context, filename string) (string, error) {
	presignClient := s3.NewPresignClient(s.Client)
	req, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (s *S3Service) GenerateDownloadURL(ctx context.Context, filename string) (string, error) {
	presignClient := s3.NewPresignClient(s.Client)
	req, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}
