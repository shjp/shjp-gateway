package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// QueryService manages query requests
type QueryService struct {
	BaseURL string
}

// NewQueryService creates a new Queryservice
func NewQueryService(baseURL string) (*QueryService, error) {
	log.Println("Initializing query service... | URL:", baseURL)
	return &QueryService{BaseURL: baseURL}, nil
}

func (s *QueryService) url(path string) string {
	return fmt.Sprintf("%s/%s", s.BaseURL, path)
}

func (s *QueryService) getOne(model, id string) (map[string]interface{}, error) {
	url := s.url(fmt.Sprintf("%s/%s", model, id))

	client := &http.Client{Timeout: time.Second * 10}
	log.Println("Sending GET request to", url)
	resp, err := client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Request to get %s failed", url))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot read response body for the request to %s", url))
	}

	var result map[string]interface{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot unmarshal response body %s", string(body)))
	}

	return result, nil
}

func (s *QueryService) getAll(model string) ([]interface{}, error) {
	url := s.url(model)

	client := &http.Client{Timeout: time.Second * 10}
	log.Println("Sending GET request to", url)
	resp, err := client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Request to get %s failed", url))
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot read response body for the request to %s", url))
	}

	var result []interface{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Cannot unmarshal response body %s", string(body)))
	}

	return result, nil
}
