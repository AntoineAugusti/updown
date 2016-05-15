package updown

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TQToken = "s7su"

func newClient() *Client {
	apiKey := os.Getenv("UPDOWN_API_KEY")
	if apiKey == "" {
		panic("API key is not set")
	}
	return NewClient(apiKey, nil)
}

func TestTokenForAlias(t *testing.T) {
	client := newClient()
	// Cache miss + alias not found
	token, err := client.Check.TokenForAlias("foo")
	assert.Equal(t, "", token)
	assert.Equal(t, ErrTokenNotFound, err)

	// - Cache miss + match found after request
	// - Cache hit
	for i := 0; i < 2; i++ {
		token, err = client.Check.TokenForAlias("Teen Quotes")
		assert.Nil(t, err)
		assert.Equal(t, TQToken, token)
	}
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

	check, resp, err := client.Check.Get("aaaaaa")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "GET https://updown.io/api/checks/aaaaaa: 404 ", err.Error())
}

func TestListDowntimes(t *testing.T) {
	client := newClient()
	// Page should be set to 1 automatically
	downs, resp, _ := client.Downtime.List(TQToken, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.True(t, len(downs) > 1)

	// Page with no downtimes
	downs, resp, _ = client.Downtime.List(TQToken, 200)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, 0, len(downs))
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
	metricRes, resp, _ := client.Metric.List(TQToken, "host", "2016-04-01 00:00:00 +0200", "2016-04-15 00:00:00 +0200")

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	locations := [4]string{"sgp", "sfo", "gra", "alpha"}
	for _, location := range locations {
		assert.Contains(t, metricRes, location)
	}
	assert.True(t, len(metricRes) > 1)
}
