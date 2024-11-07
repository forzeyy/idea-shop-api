package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3Client() *s3.S3 {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ru-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("CLOUDRU_ACCESS_KEY"),
			os.Getenv("CLOUDRU_SECRET_KEY"),
			"",
		),
		Endpoint: aws.String("..."),
	}))

	return s3.New(sess)
}

var S3Client = NewS3Client()
