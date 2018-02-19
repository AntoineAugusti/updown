package updown

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.3"
	defaultBaseURL = "https://updown.io/api/"
	userAgent      = "Go Updown v" + libraryVersion
	mediaType      = "application/json"
)

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	Message string
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}

// Client manages communication the API
type Client struct {
	// HTTP client used to communicate with the API
	client *http.Client

	// Base URL for API requests
	BaseURL *url.URL

	// User agent for client
	UserAgent string

	// APIKey to use for the API
	APIKey string

	// Services used for communications with the API
	Check    CheckService
	Downtime DowntimeService
	Metric   MetricService
	Node     NodeService
}

// NewClient returns a new API client.
func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
		APIKey:    apiKey,
	}
	c.Check = CheckService{client: c, cache: NewMemoryCache()}
	c.Downtime = DowntimeService{client: c}
	c.Metric = MetricService{client: c}
	c.Node = NodeService{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("X-API-KEY", c.APIKey)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := response.Body.Close(); err == nil {
			err = rerr
		}
	}()

	err = CheckResponse(response)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err := io.Copy(w, response.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err := json.NewDecoder(response.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err
}

// CheckResponse checks the API response for errors, and returns them if present. A response is considered an
// error if it has a status code outside the 200 range. API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			return err
		}
	}

	return errorResponse
}
