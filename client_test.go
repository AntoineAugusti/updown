package updown

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

const TQToken = "s7su"

func newClient() *Client {
	apiKey := os.Getenv("UPDOWN_API_KEY")
	if apiKey == "" {
		panic("API key is not set")
	}
	return NewClient(apiKey, nil)
}

func TestList(t *testing.T) {
	client := newClient()
	checks, resp, _ := client.Check.List()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, len(checks) > 0)
	found := false
	for _, element := range checks {
		if element.Alias == "Teen Quotes" {
			found = true
			break
		}
	}
	assert.True(t, found, "Cannot found the Teen Quotes check")
}

func TestGet(t *testing.T) {
	client := newClient()
	check, resp, _ := client.Check.Get(TQToken)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Teen Quotes", check.Alias)
}

func TestListDowntimes(t *testing.T) {
	client := newClient()
	downs, resp, _ := client.Downtime.List(TQToken, 1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, len(downs) > 1)
}

func TestAddUpdateRemoveCheck(t *testing.T) {
	client := newClient()
	res, resp, _ := client.Check.Add(CheckItem{URL: "https://google.fr"})
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "https://google.fr", res.URL)

	res, resp, _ = client.Check.Update(res.Token, CheckItem{URL: "https://google.com"})
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "https://google.com", res.URL)

	result, resp, _ := client.Check.Remove(res.Token)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, result)
}

func TestListMetrics(t *testing.T) {
	client := newClient()
	metricRes, resp, _ := client.Metric.List(TQToken, "host", "", "")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	locations := [3]string{"gra", "gra", "sfo"}
	for _, location := range locations {
		assert.Contains(t, metricRes, location)
	}
	assert.True(t, len(metricRes) > 1)
}
