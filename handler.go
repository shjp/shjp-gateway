package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

type graphQLPostBody struct {
	Query string `json:"query"`
}

func GraphqlHandler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("token")
		log.Printf("url = %+v", r.URL)

		var requestString string
		if r.Method == http.MethodGet {
			requestString = r.URL.Query().Get("query")
		} else if r.Method == http.MethodPost {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				SendErrorResponse(w, errors.New("Cannot read POST body"), 400)
				return
			}
			var queryWrapper graphQLPostBody
			if err = json.Unmarshal(body, &queryWrapper); err != nil {
				SendErrorResponse(w, errors.New("Failed parsing POST body. Did you forget query property?"), 400)
				return
			}
			requestString = queryWrapper.Query
		} else {
			SendErrorResponse(w, errors.New("Only GET and POST allowed for GraphQL request"), 400)
			return
		}
		log.Printf("query = %s", requestString)

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: requestString,
			Context:       context.WithValue(context.Background(), "token", authToken),
		})
		if len(result.Errors) > 0 {
			log.Printf("graphql errors: %v", result.Errors)
			return
		}

		fmt.Printf("data = %+v\n", result.Data)
		fmt.Printf("errors = %+v\n", result.Errors)

		respJSON, err := json.Marshal(result)
		if err != nil {
			log.Printf("Error marshaling result: %s", err)
			SendErrorResponse(w, err, 500)
			return
		}

		SendResponse(w, string(respJSON), 200)
	}
}
