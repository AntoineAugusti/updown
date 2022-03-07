// /!\ wehbooks have been deprecated in favor of recipients
// https://updown.io/api#GET-/api/webhooks
//
package updown

import (
	"fmt"
	"net/http"
)

// Webhook represents a webhook called by Updown on any event
type Webhook struct {
	ID  string `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

// WebhookService interacts with the webhooks section of the API
type WebhookService struct {
	client *Client
}

// List lists all the webhooks
func (s *WebhookService) List() ([]Webhook, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "webhooks", nil)
	if err != nil {
		return nil, nil, err
	}

	var res []Webhook
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}

// Add adds a new webhook you want to be performed
func (s *WebhookService) Add(webhook Webhook) (Webhook, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "webhooks", webhook)
	if err != nil {
		return webhook, nil, err
	}

	var webhookWithID Webhook
	resp, err := s.client.Do(req, &webhookWithID)
	if err != nil {
		return webhook, resp, err
	}

	return webhookWithID, resp, err
}

// Remove removes a webhook from Updown by its ID
func (s *WebhookService) Remove(id string) (bool, *http.Response, error) {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("webhooks/%s", id), nil)
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
