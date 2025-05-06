package providers

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/wellmtx/challenge/internal/infra/utils"
)

type RecognitionProvider interface {
	CompareFaces(sourceImage, targetImage []byte) (bool, error)
}

type recognitionProvider struct {
	provider *rekognition.Client
}

func NewRecognitionProvider() RecognitionProvider {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(fmt.Errorf("unable to load SDK config, %v", err))
	}
	rekognitionClient := rekognition.NewFromConfig(cfg)

	return &recognitionProvider{
		provider: rekognitionClient,
	}
}

func (r *recognitionProvider) CompareFaces(sourceImage, targetImage []byte) (bool, error) {
	input := &rekognition.CompareFacesInput{
		SourceImage: &types.Image{
			Bytes: sourceImage,
		},
		TargetImage: &types.Image{
			Bytes: targetImage,
		},
		SimilarityThreshold: utils.AwsFloat32(90.0),
	}

	output, err := r.provider.CompareFaces(context.TODO(), input)
	if err != nil {
		return false, fmt.Errorf("unable to compare faces, %v", err)
	}

	return len(output.FaceMatches) > 0, nil
}
