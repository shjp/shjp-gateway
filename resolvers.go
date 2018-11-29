package gateway

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
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
		var params Params
		switch typ {
		case "announcement":
			params = &announcement{}
		case "event":
			params = &event{}
		case "group":
			params = &group{}
		case "user":
			params = &user{}
		default:
			log.Println("Model type not recognized:", typ)
		}

		if err := params.ReadParams(p); err != nil {
			log.Println(err)
		}

		params.GenerateID()

		resp, err := s.mutateModel(typ, params)
		if err != nil {
			log.Println(err)
		}

		return resp, err
	}
}

func updateModelResolver(s *MutationService, typ string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		var params Params
		switch typ {
		case "announcement":
			params = &announcement{}
		case "event":
			params = &event{}
		case "group":
			params = &group{}
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

		resp, err := s.mutateModel(typ, params)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return resp, err
	}
}

func login(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil
	/*if p.Args["accountId"] == nil {
		return nil, errors.New("accountId must be present")
	}
	accountID := p.Args["accountId"].(string)

	if p.Args["clientId"] == nil {
		return nil, errors.New("clientId must be present")
	}
	clientID := p.Args["clientId"].(string)

	if p.Args["accountType"] == nil {
		return nil, errors.New("account type must be present")
	}
	loginType := p.Args["accountType"].(string)

	var profileImage *string
	if p.Args["profileImage"] == nil {
		profileImage = nil
	} else {
		*profileImage = p.Args["profileImage"].(string)
	}

	var nickname *string
	if p.Args["nickname"] == nil {
		nickname = nil
	} else {
		*nickname = p.Args["nickname"].(string)
	}

	m := models.NewMember()
	switch loginType {
	case "email":
		m.AccountType = constant.Email
	case "kakao":
		m.AccountType = constant.Kakao
	case "facebook":
		m.AccountType = constant.Facebook
	case "gmail":
		m.AccountType = constant.Gmail
	default:
		m.AccountType = constant.Undefined
	}
	if m.AccountType == constant.Undefined {
		return nil, fmt.Errorf("Cannot recognize account type %s", loginType)
	}

	return m.Login(accountID, clientID, profileImage, nickname)*/
}
