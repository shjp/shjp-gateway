package gateway

import (
	"log"

	"github.com/graphql-go/graphql"
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
