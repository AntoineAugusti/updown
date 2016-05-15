package updown

import (
	"sync"
)

// Cache lets you cache indefinitely values
type Cache interface {
	Has(key string) bool
	Put(key, value string)
	Get(key string) (has bool, value string)
}

// MemoryCache is a cache that works in memory
type MemoryCache struct {
	items map[string]string
	mu    sync.RWMutex
}

// Has determines if we can find in the cache a key for the given value
func (c *MemoryCache) Has(key string) bool {
	c.mu.RLock()
	_, has := c.items[key]
	c.mu.RUnlock()
	return has
}

// Put associates a key to a given value in the cache
func (c *MemoryCache) Put(key, value string) {
	c.mu.Lock()
	c.items[key] = value
	c.mu.Unlock()
}

// Get gets a value from the cache by its key and tells if it was found or not
func (c *MemoryCache) Get(key string) (has bool, value string) {
	c.mu.RLock()
	value, has = c.items[key]
	c.mu.RUnlock()
	return
}

// NewMemoryCache creates a new memory cache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{items: make(map[string]string)}
}
