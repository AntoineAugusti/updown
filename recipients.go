package updown

import (
	"fmt"
	"net/http"
)

// Recipients represents a end point type called by Updown on any event
type Recipient struct {
	ID    string `json:"id,omitempty"`
	TYPE  string `json:"type,omitempty"`
	VALUE string `json:"value,omitempty"`
}

// CheckService interacts with the checks section of the API
type RecipientService struct {
	client *Client
	cache  Cache
}

// List lists all the recipients
func (s *RecipientsService) List() ([]Recipients, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "recipients", nil)
	if err != nil {
		return nil, nil, err
	}

	var res []Recipients
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}

// Add adds a new recipients you want to be performed
func (s *RecipientsService) Add(data Recipients) (Recipients, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "recipients", data)
	if err != nil {
		return Recipients, nil, err
	}

	var res Recipients
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return Check{}, resp, err
	}

	return res, resp, err
}

// Remove removes a recipients from Updown by its ID
func (s *RecipientsService) Remove(id string) (bool, *http.Response, error) {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("recipients/%s", id), nil)
	if err != nil {
		return false, nil, err
	}

	var res struct {
		Deleted bool `json:"deleted,omitempty"`
	}

	resp, err := s.client.Do(req, &res)
	if err != nil {
		return false, resp, err
	}

	return res.Deleted, resp, err
}
