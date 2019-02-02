package gateway

import (
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

func transformTypeFieldsToArgument(o graphql.Object, fields ...string) graphql.FieldConfigArgument {
	a := make(graphql.FieldConfigArgument)
	for _, name := range fields {
		field, ok := o.Fields()[name]
		if !ok {
			log.Println("Undefined GraphQL field given: ", name)
		} else {
			a[name] = &graphql.ArgumentConfig{Type: field.Type}
		}
	}
	return a
}

func cleanseReturnObject(input interface{}) (map[string]interface{}, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "Error marshaling input")
	}

	var ret map[string]interface{}
	if err = json.Unmarshal(bytes, &ret); err != nil {
		return nil, errors.Wrap(err, "Error unmarshaling bytes")
	}

	return ret, nil
}
