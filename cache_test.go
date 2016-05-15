package updown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	c := NewMemoryCache()

	assert.False(t, c.Has("foo"))

	c.Put("foo", "bar")
	assert.True(t, c.Has("foo"))
	has, val := c.Get("foo")
	assert.True(t, has)
	assert.Equal(t, "bar", val)
}
