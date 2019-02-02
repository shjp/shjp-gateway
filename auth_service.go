package gateway

import (
	"encoding/json"
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

// Authenticate performs a simple authentication with the access token
func (s *AuthService) Authenticate(token string) (*auth.UserSession, error) {
	sessionClient, err := s.sessionClientOptions.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "Error creating the session client")
	}

	raw, err := sessionClient.Get(token)
	if err != nil {
		return nil, errors.Wrap(err, "Error retrieving user session")
	}

	if raw == nil {
		return nil, errors.Wrap(err, "User session doesn't exist")
	}

	var u core.User
	if err = json.Unmarshal(raw, &u); err != nil {
		return nil, errors.Wrap(err, "Error unmarshaling to user object")
	}

	return &auth.UserSession{
		Key:  token,
		User: u,
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
