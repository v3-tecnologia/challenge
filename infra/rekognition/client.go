package rekognition

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
)

var Client *rekognition.Client

func InitRekognitionClient(ctx context.Context) (*rekognition.Client, error) {
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_REK_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_REK_SECRET_ACCESS_KEY")
	sessionToken := os.Getenv("AWS_SESSION_TOKEN")

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey, secretKey, sessionToken,
		)),
	)
	if err != nil {
		return nil, err
	}

	return rekognition.NewFromConfig(cfg), nil
}
