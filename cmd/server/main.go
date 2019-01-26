package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/joho/godotenv"

	"github.com/shjp/shjp-auth/redis"
	"github.com/shjp/shjp-gateway"
)

func main() {
	envVars, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	daoHost := envVars["DAO_URL"]
	queueHost := envVars["QUEUE_URL"]
	queueUser := envVars["QUEUE_USER"]
	queueExchange := envVars["QUEUE_EXCHANGE"]
	redisHost := envVars["REDIS_URL"]

	queryService, err := gateway.NewQueryService(daoHost)
	if err != nil {
		log.Fatalf("Failed instantiating a query service: %s", err)
		return
	}
	mutationService, err := gateway.NewMutationService(queueHost, queueUser, queueExchange)
	if err != nil {
		log.Fatalf("Failed instantiating a mutation service: %s", err)
		return
	}
	authService, err := gateway.NewAuthService(daoHost+"/users/search", &redis.Options{
		Network: "tcp",
		Address: redisHost,
	})
	if err != nil {
		log.Fatalf("Failed instantiating an auth service: %s", err)
		return
	}
	schema, err := gateway.ConfigSchema(queryService, mutationService, authService)
	if err != nil {
		log.Fatalf("Failed configuring schema: %v", err)
		return
	}

	// Below handler is for dev purpose
	interactiveGqHandler := handler.New(&handler.Config{Schema: &schema, Pretty: true, GraphiQL: true})
	http.Handle("/console", interactiveGqHandler)

	http.HandleFunc("/graphql", gateway.GraphqlHandler(schema))

	log.Println("Server listening to port 8100") //
	log.Fatal(http.ListenAndServe(":8100", nil))
}
