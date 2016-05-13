package updown

import (
	"fmt"
	"net/http"
)

type Ssl struct {
	TestedAt string `json:"tested_at,omitempty"`
	Valid    bool   `json:"valid,omitempty"`
	Error    string `json:"error,omitempty"`
}

type Check struct {
	Token       string  `json:"token,omitempty"`
	Url         string  `json:"url,omitempty"`
	Alias       string  `json:"alias,omitempty"`
	LastStatus  int     `json:"last_status,omitempty"`
	Uptime      float64 `json:"uptime,omitempty"`
	Down        bool    `json:"down,omitempty"`
	DownSince   string  `json:"down_since,omitempty"`
	Error       string  `json:"error,omitempty"`
	Period      int     `json:"period,omitempty"`
	Apdex_t     float64 `json:"apdex_t,omitempty"`
	Enabled     bool    `json:"enabled,omitempty"`
	Published   bool    `json:"published,omitempty"`
	LastCheckAt string  `json:"last_check_at,omitempty"`
	NextCheckAt string  `json:"next_check_at,omitempty"`
	FaviconUrl  string  `json:"favicon_url,omitempty"`
	Ssl         Ssl     `json:"ssl,omitempty"`
}

type CheckItem struct {
	// The URL you want to monitor
	Url string `json:"url,omitempty"`
	// Interval in seconds (30, 60, 120, 300 or 600)
	Period int `json:period,omitempty`
	// APDEX threshold in seconds (0.125, 0.25, 0.5 or 1.0)
	Apdex float64 `json:apdex,omitempty`
	// Is the check enabled
	Enabled bool `json:enabled,omitempty`
	// Shall the status page be public
	Published bool `json:published,omitempty`
	// Human readable name
	Alias string `json:alias,omitempty`
	// Search for this string in the page
	StringMatch string `json:stringmatch,omitempty`
}

type CheckService struct {
	client *Client
}

type RemoveResponse struct {
	Deleted bool `json:"deleted,omitempty"`
}

func (s *CheckService) List() ([]Check, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "checks", nil)
	if err != nil {
		return nil, nil, err
	}

	var res []Check
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}

func (s *CheckService) Get(token string) (Check, *http.Response, error) {
	req, err := s.client.NewRequest("GET", pathForToken(token), nil)
	if err != nil {
		return Check{}, nil, err
	}

	var res Check
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return Check{}, resp, err
	}

	return res, resp, err
}

func (s *CheckService) Add(data CheckItem) (Check, *http.Response, error) {
	req, err := s.client.NewRequest("POST", "checks", data)
	if err != nil {
		return Check{}, nil, err
	}

	var res Check
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return Check{}, resp, err
	}

	return res, resp, err
}

func (s *CheckService) Update(token string, data CheckItem) (Check, *http.Response, error) {
	req, err := s.client.NewRequest("PUT", pathForToken(token), data)
	if err != nil {
		return Check{}, nil, err
	}

	var res Check
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return Check{}, resp, err
	}

	return res, resp, err
}

func (s *CheckService) Remove(token string) (RemoveResponse, *http.Response, error) {
	req, err := s.client.NewRequest("DELETE", pathForToken(token), nil)
	if err != nil {
		return RemoveResponse{Deleted: false}, nil, err
	}

	var res RemoveResponse
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return RemoveResponse{Deleted: false}, resp, err
	}

	return res, resp, err
}

func pathForToken(token string) string {
	return fmt.Sprintf("checks/%s", token)
}
