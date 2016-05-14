package updown

import (
	"fmt"
	"net/http"
)

type ResponseTime struct {
	Under125  int `json:"under125,omitempty"`
	Under250  int `json:"under250,omitempty"`
	Under500  int `json:"under500,omitempty"`
	Under1000 int `json:"under1000,omitempty"`
	Under2000 int `json:"under2000,omitempty"`
	Under4000 int `json:"under4000,omitempty"`
}

type Requests struct {
	Samples      int          `json:"samples,omitempty"`
	Failures     int          `json:"failures,omitempty"`
	Satisfied    int          `json:"satisfied,omitempty"`
	Tolerated    int          `json:"tolerated,omitempty"`
	ResponseTime ResponseTime `json:"by_response_time,omitempty"`
}

type Host struct {
	Ip          string `json:"ip,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type Timings struct {
	Redirect   int `json:"redirect,omitempty"`
	NameLookup int `json:"namelookup,omitempty"`
	Connection int `json:"connection,omitempty"`
	Handshake  int `json:"handshake,omitempty"`
	Response   int `json:"response,omitempty"`
	Total      int `json:"total,omitempty"`
}

type MetricItem struct {
	Apdex    float64  `json:"apdex,omitempty"`
	Requests Requests `json:"requests,omitempty"`
	Timings  Timings  `json:"timings,omitempty"`
	Host     Host     `json:"host,omitempty"`
}

type Metrics map[string]MetricItem

type MetricService struct {
	client *Client
}

func (s *MetricService) List(token, group, from, to string) (Metrics, *http.Response, error) {
	path := fmt.Sprintf(pathForToken(token)+"/metrics?group=%s", group)
	// Optional from and to parameters
	if from != "" {
		path += "&from=" + from
	}
	if to != "" {
		path += "&to=" + to
	}
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var res Metrics
	resp, err := s.client.Do(req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res, resp, err
}
