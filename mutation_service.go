package gateway

import (
	"log"

	"github.com/pkg/errors"

	"github.com/shjp/shjp-core"
	"github.com/shjp/shjp-queue"
)

// MutationService manages mutation requests
type MutationService struct {
	producer *queue.Producer
	exchange string
}

// MutationResponse specifies the format of the mutations responses
type MutationResponse struct {
	RefID string `json:"ref_id"`
}

// NewMutationService creates a new MutationService
func NewMutationService(baseURL, user, exchange string) (*MutationService, error) {
	log.Println("Initializing mutation service... | URL:", baseURL, "| User:", user, "| Exchange:", exchange)
	producer, err := queue.NewProducer(baseURL, user)
	if err != nil {
		return nil, errors.Wrap(err, "Error instantiating a producer")
	}
	return &MutationService{producer, exchange}, nil
}

func (s *MutationService) dispatch(m *core.Message) error {
	log.Printf("Dispatching message for mutations: %+v\n", *m)
	return s.producer.Publish(s.exchange, m)
}

func (s *MutationService) mutateModel(typ string, p Params) (*MutationResponse, error) {
	msg, err := p.Pack(core.IntentRequest, core.UpsertOperation)
	resp := &MutationResponse{RefID: msg.Key}
	if err != nil {
		return resp, errors.Wrap(err, "Error packing the params as message")
	}

	if err = s.dispatch(msg); err != nil {
		return resp, errors.Wrap(err, "Error dispatching mutation message")
	}

	return resp, nil
}
