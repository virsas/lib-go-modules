package vsslib

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type DynamoHandler interface {
	List(filter expression.ConditionBuilder, projection expression.ProjectionBuilder) ([]map[string]interface{}, error)
	Create(map[string]interface{}) error
	Show(key map[string]*dynamodb.AttributeValue) (map[string]interface{}, error)
	Delete(key map[string]*dynamodb.AttributeValue) error
}

type dynamo struct {
	session *dynamodb.DynamoDB
	table   string
}

func NewDynamoSession(sess *session.Session) (DynamoHandler, error) {
	var dynamoTable string = ""
	dynamoTableValue, dynamoTablePresent := os.LookupEnv("AWS_DYNAMO_TABLE")
	if dynamoTablePresent {
		dynamoTable = dynamoTableValue
	} else {
		panic("Missing ENV Variable AWS_DYNAMO_TABLE")
	}

	ddb, err := newDynamoSession(sess, dynamoTable)
	if err != nil {
		return nil, err
	}

	return ddb, nil
}

func newDynamoSession(sess *session.Session, table string) (DynamoHandler, error) {
	ddbsess := dynamodb.New(sess)
	ddb := &dynamo{session: ddbsess, table: table}

	return ddb, nil
}

func (d *dynamo) List(filter expression.ConditionBuilder, projection expression.ProjectionBuilder) ([]map[string]interface{}, error) {
	var err error

	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(projection).Build()
	if err != nil {
		return nil, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(d.table),
	}

	result, err := d.session.Scan(params)
	if err != nil {
		return nil, err
	}

	items := []map[string]interface{}{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (d *dynamo) Create(item map[string]interface{}) error {
	var err error

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(d.table),
	}

	_, err = d.session.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (d *dynamo) Show(key map[string]*dynamodb.AttributeValue) (map[string]interface{}, error) {
	var err error
	var item map[string]interface{} = map[string]interface{}{}

	result, err := d.session.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(d.table),
		Key:       key,
	})
	if err != nil {
		return item, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (d *dynamo) Delete(key map[string]*dynamodb.AttributeValue) error {
	var err error

	input := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: aws.String(d.table),
	}

	_, err = d.session.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}
