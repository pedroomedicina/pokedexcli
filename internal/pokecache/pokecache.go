package pokecache

import (
	"fmt"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(reapInterval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}

	go c.reapLoop(reapInterval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}

	fmt.Println("Added key:", key)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheE, ok := c.cache[key]

	return cacheE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for key, val := range c.cache {
		if val.createdAt.Before(timeAgo) {
			delete(c.cache, key)
		}
	}
}
