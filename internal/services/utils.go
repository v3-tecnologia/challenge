package services

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

var (
	s3Region           string = os.Getenv("AWS_REGION")
	awsAccessKeyID     string = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey string = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsConfig                 = aws.Config{
		Region:      s3Region,
		Credentials: credentials.NewStaticCredentialsProvider(awsAccessKeyID, awsSecretAccessKey, ""),
	}
)
