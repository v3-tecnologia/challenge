package rekognition

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

type RekognitionClient struct {
	client *rekognition.Client
}

func NewClient(ctx context.Context) (*RekognitionClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao carregar config AWS: %w", err)
	}
	return &RekognitionClient{
		client: rekognition.NewFromConfig(cfg),
	}, nil
}

func (r *RekognitionClient) DetectFaces(imageBytes []byte) (*rekognition.DetectFacesOutput, error) {
	input := &rekognition.DetectFacesInput{
		Image: &types.Image{
			Bytes: imageBytes,
		},
		Attributes: []types.Attribute{types.AttributeAll},
	}
	return r.client.DetectFaces(context.TODO(), input)
}
