package cache

import (
    "sync"
)

type Cache[K comparable, V any] struct {
    items map[K]V
    mu    sync.Mutex
}

// Cache[Username, AuthToken]
var AuthCache *Cache[string, string]

func Init[K string, V string]() {
    AuthCache = &Cache[string, string]{
        items: make(map[string]string),
    }
}

func (c *Cache[K, V]) Set(key K, value V) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.items[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    value, found := c.items[key]
    return value, found
}

func (c *Cache[K, V]) Remove(key K) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    delete(c.items, key)
}

