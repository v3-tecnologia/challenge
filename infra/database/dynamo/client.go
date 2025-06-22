package dynamo

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var Client *dynamodb.Client

func InitDynamoClient(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, _ ...interface{}) (aws.Endpoint, error) {
				if service == dynamodb.ServiceID {
					return aws.Endpoint{
						URL:           os.Getenv("DYNAMO_ENDPOINT"),
						SigningRegion: "us-east-1",
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			}),
		),
	)
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}
