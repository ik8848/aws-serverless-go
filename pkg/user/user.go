package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/ikatwal/aws-serverless-go/pkg/validators"
)

var (
	ErrorInvalidUserData         = "invalid user data"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorInvalidEmail            = "invalid email"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotPutItemInDB     = "could not put item in db"
	ErrorUserAlreadyExists       = "user already exists"
	ErrorUserDoesNotExist        = "user does not exist"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetUser(email, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}
	result, err := dbClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	user := &User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, user)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return user, nil
}

func GetUsers(tableName string, dbClient dynamodbiface.DynamoDBAPI) ([]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := dbClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	users := make([]User, 0)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return users, nil
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var u User
	if err := json.Unmarshal([]byte(req.Body), &u); err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	if !validators.IsValidEmail(u.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}
	user, _ := GetUser(u.Email, tableName, dbClient)
	if user != nil && len(user.Email) != 0 {
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	jsonUser, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}
	input := &dynamodb.PutItemInput{
		Item:      jsonUser,
		TableName: aws.String(tableName),
	}
	_, err = dbClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotPutItemInDB)
	}
	return &u, nil
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

}
