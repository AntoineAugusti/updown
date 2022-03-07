package updown

import (
	"fmt"
	"net/http"
)

// Recipient represents a Recipient called by Updown on any event
type Recipient struct {
	ID        string        `json:"id,omitempty"`
	Type      RecipientType `json:"type,omitempty"`
	Name      string        `json:"name,omitempty"`
	Immutable bool          `json:"immutable,omitempty"`
}

type RecipientType string

const (
	RecipientTypeEmail    RecipientType = "email"
	RecipientTypeSMS      RecipientType = "sms"
	RecipientTypeTelegram RecipientType = "telegram"
	RecipientTypeWebhook  RecipientType = "webhook"
	RecipientTypeZapier   RecipientType = "zapier"
)

// RecipientService interacts with the Recipients section of the API
type RecipientService struct {
	client *Client
}

// List lists all the Recipients
func (s *RecipientService) List() ([]Recipient, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "recipients", nil)
	if err != nil {
		return nil, nil, err
	}

	var res []Recipient
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}

// Add adds a new Recipient you want to be performed
func (s *RecipientService) Add(recipient Recipient) (Recipient, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "recipients", recipient)
	if err != nil {
		return recipient, nil, err
	}

	var recipientWithID Recipient
	resp, err := s.client.Do(req, &recipientWithID)
	if err != nil {
		return recipient, resp, err
	}

	return recipientWithID, resp, err
}

// Remove removes a Recipient from Updown by its ID
func (s *RecipientService) Remove(id string) (bool, *http.Response, error) {
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
