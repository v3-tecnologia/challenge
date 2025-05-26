package s3

import (
	"context"
	"fmt"
	"log"

	env "github.com/KaiRibeiro/challenge/internal/config"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func InitS3() {
	ctx := context.Background()
	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	S3Client = s3.NewFromConfig(awsCfg)
	fmt.Println("Connected to S3")
}
