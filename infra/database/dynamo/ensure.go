package dynamo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func EnsureTable(ctx context.Context, client *dynamodb.Client, tableName string) error {
	_, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: &tableName,
	})

	if err == nil {
		log.Printf("Tabela %s já existe", tableName)
		return nil
	}

	var rnfe *types.ResourceNotFoundException
	if !errors.As(err, &rnfe) {
		return fmt.Errorf("erro ao verificar tabela: %w", err)
	}

	_, err = client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: &tableName,
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("uuid"),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("uuid"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})

	if err != nil {
		return fmt.Errorf("erro ao criar tabela: %w", err)
	}

	log.Printf("Tabela %s criada com sucesso, aguardando estar ativa...", tableName)

	for {
		out, err := client.DescribeTable(ctx, &dynamodb.DescribeTableInput{
			TableName: &tableName,
		})
		if err != nil {
			return fmt.Errorf("erro ao verificar status da tabela: %w", err)
		}

		if out.Table.TableStatus == types.TableStatusActive {
			break
		}

		log.Printf("Aguardando tabela %s ficar ativa...", tableName)
		time.Sleep(2 * time.Second)
	}

	return nil
}
