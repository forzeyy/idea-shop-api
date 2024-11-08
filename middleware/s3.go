package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Service struct {
	Client *s3.Client
	Bucket string
}

func NewS3Service(client *s3.Client, bucket string) *S3Service {
	return &S3Service{Client: client, Bucket: bucket}
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

func (s *S3Service) UploadImageToS3(ctx context.Context, productID uint, file multipart.File, filename string) (string, error) {

	key := fmt.Sprintf("%v-%v", productID, filename)

	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", fmt.Errorf("failed to copy file to buffer: %v", err)
	}

	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String("image/jpeg"),
		ACL:         types.ObjectCannedACLPublicRead,
	}

	_, err := s.Client.PutObject(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	fileURL := fmt.Sprintf("https://%s.storage.yandexcloud.net/%s", s.Bucket, key)
	return fileURL, nil
}
