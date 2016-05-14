package updown

import (
	"net/http"
	"net/url"
)

// ResponseTime represents the response times in milliseconds
type ResponseTime struct {
	Under125  int `json:"under125,omitempty"`
	Under250  int `json:"under250,omitempty"`
	Under500  int `json:"under500,omitempty"`
	Under1000 int `json:"under1000,omitempty"`
	Under2000 int `json:"under2000,omitempty"`
	Under4000 int `json:"under4000,omitempty"`
}

// Requests gives statistics about requests made to check the status
type Requests struct {
	Samples      int          `json:"samples,omitempty"`
	Failures     int          `json:"failures,omitempty"`
	Satisfied    int          `json:"satisfied,omitempty"`
	Tolerated    int          `json:"tolerated,omitempty"`
	ResponseTime ResponseTime `json:"by_response_time,omitempty"`
}

// Host represents the host where the check was made
type Host struct {
	IP          string `json:"ip,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// Timings represents the amount of time taken by each part of the connection
type Timings struct {
	Redirect   int `json:"redirect,omitempty"`
	NameLookup int `json:"namelookup,omitempty"`
	Connection int `json:"connection,omitempty"`
	Handshake  int `json:"handshake,omitempty"`
	Response   int `json:"response,omitempty"`
	Total      int `json:"total,omitempty"`
}

// MetricItem is basically the core metric
type MetricItem struct {
	Apdex    float64  `json:"apdex,omitempty"`
	Requests Requests `json:"requests,omitempty"`
	Timings  Timings  `json:"timings,omitempty"`
	Host     Host     `json:"host,omitempty"`
}

// Metrics represents multiple metrics
type Metrics map[string]MetricItem

// MetricService interacts with the metrics section of the API
type MetricService struct {
	client *Client
}

// List lists metrics available for a check identified by a taken, grouped by the given group
// (host|time) over a period
func (s *MetricService) List(token, group, from, to string) (Metrics, *http.Response, error) {
	u, _ := url.Parse(pathForToken(token) + "/metrics")
	q := u.Query()
	q.Add("group", group)

	// Optional from and to parameters
	if from != "" {
		q.Add("from", from)
	}
	if to != "" {
		q.Add("to", to)
	}
	u.RawQuery = q.Encode()

	req, err := s.client.NewRequest("GET", u.String(), nil)
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
