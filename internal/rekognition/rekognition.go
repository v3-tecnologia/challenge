package rekognition

import (
	"context"
	"fmt"

	env "github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/logs"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
)

var RekognitionClient *rekognition.Client

func InitRekognition() {
	logs.Logger.Info("Creating Rekognition client")
	ctx := context.Background()
	AwsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))

	if err != nil {
		wrappedErr := fmt.Errorf("failed to load AWS Rekognition config: %w", err)
		logs.Logger.Error("failed to load AWS config",
			"error", wrappedErr,
		)
	}

	RekognitionClient = rekognition.NewFromConfig(AwsCfg)
	logs.Logger.Info("Rekognition client created")
}
