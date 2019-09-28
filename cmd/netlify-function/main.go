package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
	gateway "github.com/sanggonlee/shjp-gateway"
	"github.com/shjp/shjp-auth/redis"
)

type contextKey string

const (
	authTokenKey contextKey = "accessToken"
)

type graphQLPostBody struct {
	Query string `json:"query"`
}

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	authToken, ok := request.Headers["Auth-Token"]
	// For time being, simply log and pass an empty string when auth token is not found
	if !ok {
		log.Println("Auth token not found")
	}

	daoHost := os.Getenv("DAO_URL")
	queueHost := os.Getenv("QUEUE_URL")
	queueUser := os.Getenv("QUEUE_USER")
	queueExchange := os.Getenv("QUEUE_EXCHANGE")
	redisHost := os.Getenv("REDIS_URL")

	queryService, err := gateway.NewQueryService(daoHost)
	if err != nil {
		log.Fatalf("Failed instantiating a query service: %s", err)
		return formatResponse(http.StatusInternalServerError, "Error init query service"), err
	}
	mutationService, err := gateway.NewMutationService(queueHost, queueUser, queueExchange)
	if err != nil {
		log.Fatalf("Failed instantiating a mutation service: %s", err)
		return formatResponse(http.StatusInternalServerError, "Error init mutation service"), err
	}
	authService, err := gateway.NewAuthService(daoHost+"/users/search", &redis.Options{
		Network: "tcp",
		Address: redisHost,
	})
	if err != nil {
		log.Fatalf("Failed instantiating an auth service: %s", err)
		return formatResponse(http.StatusInternalServerError, "Error init auth service"), err
	}
	schema, err := gateway.ConfigSchema(queryService, mutationService, authService)
	if err != nil {
		log.Fatalf("Failed configuring schema: %v", err)
		return formatResponse(http.StatusInternalServerError, "Error init GraphQL schema"), err
	}

	var requestString string
	if request.HTTPMethod == http.MethodOptions {
		return formatResponse(http.StatusOK, "ok"), nil
	}

	if request.HTTPMethod == http.MethodGet {
		var ok bool
		requestString, ok = request.QueryStringParameters["query"]
		if !ok {
			return formatResponse(http.StatusBadRequest, ""), errors.New("No query given in query params")
		}
	} else if request.HTTPMethod == http.MethodPost {
		var queryWrapper graphQLPostBody
		if err = json.Unmarshal([]byte(request.Body), &queryWrapper); err != nil {
			return formatResponse(http.StatusBadRequest, ""), errors.New("Failed parsing POST body. Did you forget query property?")
		}
		requestString = queryWrapper.Query
	} else {
		return formatResponse(http.StatusBadRequest, ""), errors.New("Only GET, POST and OPTIONS allowed for GraphQL request")
	}
	log.Printf("query = %s", requestString)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: requestString,
		Context:       context.WithValue(context.Background(), authTokenKey, authToken),
	})
	if result.HasErrors() {
		log.Printf("graphql errors: %v", result.Errors)
		result.Data = nil
		errBytes, err := json.Marshal(result)
		if err != nil {
			return formatResponse(http.StatusInternalServerError, "Unable to marshal error"), err
		}
		return formatResponse(http.StatusInternalServerError, "Error from GraphQL"), errors.New(string(errBytes))
	}

	fmt.Printf("data = %+v\n", result.Data)

	respJSON, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error marshaling result: %s", err)
		return formatResponse(http.StatusInternalServerError, "Unable to marshal result"), err
	}

	return formatResponse(http.StatusOK, string(respJSON)), nil
}

func formatResponse(statusCode int, body string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}
