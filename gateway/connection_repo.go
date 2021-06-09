package gateway

import (
	"context"
	"lambda/domain/connection"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ConnectionRepository struct {
	client     *dynamodb.Client
	primaryKey string
	sortKey    string
	tableName  string
}

func NewConnectionRepository(client *dynamodb.Client, pkey string, tblName string) *ConnectionRepository {
	return &ConnectionRepository{
		client:     client,
		primaryKey: pkey,
		tableName:  tblName,
	}
}

func (db ConnectionRepository) Save(id connection.Id) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(db.tableName),
		Item: map[string]types.AttributeValue{
			db.primaryKey: &types.AttributeValueMemberS{id.String()},
		},
	}

	_, err := db.client.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return err
}

func (db ConnectionRepository) Delete(id connection.Id) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(db.tableName),
		Key: map[string]types.AttributeValue{
			db.primaryKey: &types.AttributeValueMemberS{id.String()},
		},
	}

	_, err := db.client.DeleteItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func (db ConnectionRepository) Find(id connection.Id) error {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(db.tableName),
		Key: map[string]types.AttributeValue{
			db.primaryKey: &types.AttributeValueMemberS{id.String()},
		},
	}

	_, err := db.client.GetItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}
