package handlers

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/ishakatwal/aws-serverless-go/pkg/user"
)

var ErrorMethodNotAllowed = "method not allowed"

func GetUser(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.GetUser(email, tableName, dbClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, errors.New(ErrorMethodNotAllowed))
		}
		return apiResponse(http.StatusOK, result)
	}
	result, err := user.GetUsers(tableName, dbClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, errors.New(ErrorMethodNotAllowed))
	}
	return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req, tableName, dbClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, err)
	}
	return apiResponse(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(req, tableName, dbClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, err)
	}
	return apiResponse(http.StatusOK, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	err := user.DeleteUser(req, tableName, dbClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, err)
	}
	return apiResponse(http.StatusOK, nil)
}

func DefaultMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, errors.New(ErrorMethodNotAllowed))
}
