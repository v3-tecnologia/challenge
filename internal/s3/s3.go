package s3

import (
	"context"
	"log"
	"net/http"

	env "github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func InitS3() {
	ctx := context.Background()
	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(env.AwsRegion))
	if err != nil {
		log.Fatal(custom_errors.NewS3Error(err, http.StatusInternalServerError))
	}

	S3Client = s3.NewFromConfig(awsCfg)
}
