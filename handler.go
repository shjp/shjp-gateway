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

type contextKey string

const (
	authTokenKey contextKey = "accessToken"
)

type graphQLPostBody struct {
	Query string `json:"query"`
}

// GraphqlHandler handles requests made on console
func GraphqlHandler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Auth-Token")
		log.Printf("url = %+v", r.URL)

		var requestString string
		if r.Method == http.MethodOptions {
			SendResponse(w, "ok", http.StatusOK)
			return
		} else if r.Method == http.MethodGet {
			requestString = r.URL.Query().Get("query")
		} else if r.Method == http.MethodPost {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				SendErrorResponse(w, errors.New("Cannot read POST body"), http.StatusBadRequest)
				return
			}
			var queryWrapper graphQLPostBody
			if err = json.Unmarshal(body, &queryWrapper); err != nil {
				SendErrorResponse(w, errors.New("Failed parsing POST body. Did you forget query property?"), http.StatusBadRequest)
				return
			}
			requestString = queryWrapper.Query
		} else {
			SendErrorResponse(w, errors.New("Only GET, POST and OPTIONS allowed for GraphQL request"), http.StatusBadRequest)
			return
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
				SendErrorResponse(w, errors.New("Error marshaling errors"), http.StatusInternalServerError)
				return
			}
			SendResponse(w, string(errBytes), http.StatusInternalServerError)
			return
		}

		fmt.Printf("data = %+v\n", result.Data)

		respJSON, err := json.Marshal(result)
		if err != nil {
			log.Printf("Error marshaling result: %s", err)
			SendErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		SendResponse(w, string(respJSON), http.StatusOK)
	}
}
