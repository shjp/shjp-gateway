package gateway

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"

	core "github.com/shjp/shjp-core"
)

func queryOneResolver(s *QueryService, typ string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return s.getOne(typ, p.Args["id"].(string))
	}
}

func queryAllResolver(s *QueryService, typ string) graphql.FieldResolveFn {
	return func(_ graphql.ResolveParams) (interface{}, error) {
		return s.getAll(typ)
	}
}

func createModelResolver(s *MutationService, typ string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		var params ModelParams
		switch typ {
		case "announcement":
			params = &announcement{}
		case "event":
			params = &event{}
		case "group":
			params = &group{}
		// case "role":
		// 	params = &role{}
		case "user":
			params = &user{}
		default:
			log.Println("Model type not recognized:", typ)
		}

		if err := params.ReadParams(p); err != nil {
			log.Println(err)
		}

		params.GenerateID()

		resp, err := s.mutateEntity(params)
		if err != nil {
			log.Println(err)
		}

		return resp, err
	}
}

func updateModelResolver(s *MutationService, typ string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		var params ModelParams
		switch typ {
		case "announcement":
			params = &announcement{}
		case "event":
			params = &event{}
		case "group":
			params = &group{}
		// case "role":
		// 	params = &role{}
		case "user":
			params = &user{}
		default:
			log.Println("Model type not recognized:", typ)
			return nil, fmt.Errorf("Model type not recognized: %s", typ)
		}

		if err := params.ReadParams(p); err != nil {
			log.Println(err)
			return nil, err
		}

		if params.GetID() == "" {
			return nil, errors.New("ID is required for update mutation")
		}

		resp, err := s.mutateEntity(params)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return resp, err
	}
}

func upsertRelationshipResolver(as *AuthService, ms *MutationService, typ string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		token := p.Context.Value(authTokenKey)
		if token == nil || token == "" {
			return nil, errors.New("Missing auth token")
		}
		_, err := as.Authenticate(token.(string))
		if err != nil {
			return nil, errors.Wrap(err, "Authentication failed")
		}

		// TODO: Authorization

		var params Params
		switch typ {
		case "group_membership":
			params = &groupMembership{}
		case "update_rsvp":
			params = &eventRSVP{}
		default:
			log.Println("Relationship type not recognized:", typ)
			return nil, fmt.Errorf("Relationship type not recognized: %s", typ)
		}

		if err := params.ReadParams(p); err != nil {
			log.Println(err)
			return nil, err
		}

		resp, err := ms.mutateEntity(params)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return resp, nil
	}
}

func login(s *AuthService) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		accountType := p.Args["accountType"].(string)

		var email string
		if p.Args["email"] != nil {
			email = p.Args["email"].(string)
		}
		var password string
		if p.Args["password"] != nil {
			password = p.Args["password"].(string)
		}

		userPayload := core.User{
			AccountType: &accountType,
			Email:       &email,
			Password:    &password,
		}

		session, err := s.Login(userPayload)
		if err != nil {
			return nil, errors.Wrap(err, "Failed logging in")
		}

		ret, err := cleanseReturnObject(session)
		if err != nil {
			return nil, errors.Wrap(err, "Error cleansing return object")
		}

		return ret, nil
	}
}

func me(s *AuthService) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		paramsBlob, err := json.Marshal(p)
		if err != nil {
			log.Println("Marshalling params failed:", err)
		}
		log.Println("ResolveParams object ---------------------------------------------------")
		log.Println(string(paramsBlob))
		log.Println("------------------------------------------------------------------------")

		token := p.Context.Value(authTokenKey)
		if token == nil || token == "" {
			return nil, errors.New("Missing auth token")
		}
		session, err := s.Authenticate(token.(string))
		if err != nil {
			return nil, errors.Wrap(err, "Authentication failed")
		}

		ret, err := cleanseReturnObject(&session.User)
		if err != nil {
			return nil, errors.Wrap(err, "Error cleansing return object")
		}

		return ret, nil
	}
}
