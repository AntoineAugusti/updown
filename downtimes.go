package updown

import (
	"fmt"
	"net/http"
)

type Downtime struct {
	Error     string `json:"error,omitempty"`
	StartedAt string `json:"started_at,omitempty"`
	EndedAt   string `json:"ended_at,omitempty"`
	Duration  int    `json:"duration,omitempty"`
}

type DowntimeService struct {
	client *Client
}

func (s *DowntimeService) List(token string) ([]Downtime, *http.Response, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("checks/%s/downtimes", token), nil)
	if err != nil {
		return nil, nil, err
	}

	var res []Downtime
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}
