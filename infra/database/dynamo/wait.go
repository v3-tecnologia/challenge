package dynamo

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func WaitForDynamoReady(ctx context.Context, client *dynamodb.Client) error {
	for i := 1; i <= 10; i++ {
		_, err := client.ListTables(ctx, &dynamodb.ListTablesInput{})
		if err == nil {
			fmt.Println("DynamoDB está pronto!")
			return nil
		}

		fmt.Printf("Tentativa %d: DynamoDB ainda não está pronto (%v)\n", i, err)
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("DynamoDB não respondeu após 10 tentativas")
}
