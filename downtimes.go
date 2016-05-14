package updown

import (
	"fmt"
	"net/http"
	"strconv"
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

func (s *DowntimeService) List(token string, pageNb int) ([]Downtime, *http.Response, error) {
	path := fmt.Sprintf("checks/%s/downtimes?page=%s", token, strconv.Itoa(max(1, pageNb)))
	req, err := s.client.NewRequest("GET", path, nil)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
