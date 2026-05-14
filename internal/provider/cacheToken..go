package provider

import (
	"sync"
	"time"
)

type cachedToken struct {
	token     string
	expiresAt time.Time
}

type TokenCache struct {
	tokens map[string]*cachedToken
	mu     sync.RWMutex
}

var TokenCacheInstance = &TokenCache{
	tokens: make(map[string]*cachedToken),
}

func (c *TokenCache) Get(userID string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if ct, ok := c.tokens[userID]; ok {
		if ct.token != "" && time.Now().Before(ct.expiresAt) {
			return ct.token, true
		}
	}
	return "", false
}

func (c *TokenCache) Set(userID, token string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.tokens[userID] = &cachedToken{
		token:     token,
		expiresAt: time.Now().Add(ttl),
	}
}
