package gateway

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/shjp/shjp-auth"
	"github.com/shjp/shjp-auth/email"
	core "github.com/shjp/shjp-core"
)

// AuthService manages auth requests
type AuthService struct {
	daoURL               string
	sessionClientOptions auth.SessionClientOptions
}

// NewAuthService instantiates a new auth service
func NewAuthService(daoURL string, sessionClientOptions auth.SessionClientOptions) (*AuthService, error) {
	log.Println("Initializing auth service... | daoURL:", daoURL, "| sessionClientOptions:", sessionClientOptions.String())
	return &AuthService{
		daoURL:               daoURL,
		sessionClientOptions: sessionClientOptions,
	}, nil
}

// Login runs the login workflow
func (s *AuthService) Login(u core.User) (*auth.UserSession, error) {
	strategy, err := s.getLoginStrategy(u)
	if err != nil {
		return nil, errors.Wrap(err, "Error choosing login strategy")
	}

	sessionClient, err := s.sessionClientOptions.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "Error creating the session client")
	}

	ul := auth.NewLogin(u, strategy, sessionClient)

	session, err := ul.Login()
	if err != nil {
		return nil, errors.Wrap(err, "Failed logging in with the UserLogin instance")
	}

	return session, nil
}

func (s *AuthService) getLoginStrategy(u core.User) (auth.LoginStrategy, error) {
	if u.AccountType == nil {
		return nil, errors.New("Account type must be specified to choose login strategy")
	}

	switch *u.AccountType {
	case "email":
		return &email.LoginStrategy{DaoURL: s.daoURL}, nil
	default:
		return nil, fmt.Errorf("Unknown login type: %s", *u.AccountType)
	}
}
