package dynamo

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateTable(ctx context.Context, client *dynamodb.Client, tableName, partitionKey, sortKey string) error {
	_, err := client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String(partitionKey),
				KeyType:       types.KeyTypeHash, // Partition key
			},
			{
				AttributeName: aws.String(sortKey),
				KeyType:       types.KeyTypeRange, // Sort key
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String(partitionKey),
				AttributeType: types.ScalarAttributeTypeS, // string
			},
			{
				AttributeName: aws.String(sortKey),
				AttributeType: types.ScalarAttributeTypeS, // string
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		return fmt.Errorf("failed to create table %s: %w", tableName, err)
	}
	log.Printf("Table %s created successfully", tableName)
	return nil
}
