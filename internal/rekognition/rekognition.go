package rekognition

import (
	"context"
	"fmt"
	"log"

	env "github.com/KaiRibeiro/challenge/internal/config"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
)

var RekognitionClient *rekognition.Client

func InitRekognition() {
	ctx := context.Background()
	AwsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	RekognitionClient = rekognition.NewFromConfig(AwsCfg)
	fmt.Println("Connected to S3")
}
