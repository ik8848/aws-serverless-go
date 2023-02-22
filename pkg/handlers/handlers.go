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

}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

}

func DefaultMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, errors.New(ErrorMethodNotAllowed))
}
