package s3

import (
	"context"
	"net/http"

	env "github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/logs"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func InitS3() {
	logs.Logger.Info("starting s3 connection")
	ctx := context.Background()
	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
	if err != nil {
		wrappedErr := custom_errors.NewS3Error(err, http.StatusInternalServerError)
		logs.Logger.Error("failed to load AWS config",
			"error", wrappedErr,
		)
	}

	S3Client = s3.NewFromConfig(awsCfg)
	logs.Logger.Info("s3 is connected")
}
