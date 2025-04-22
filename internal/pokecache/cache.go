package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(d time.Duration) Cache {
	c := Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}
	go c.reapLoop(d)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{time.Now(), val}
	fmt.Printf("Cache Added: Key: %s, Length: %d\n", key, len(val))
}

func (c *Cache) Get(key string) (val []byte, b bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	result, found := c.data[key]
	if !found {
		return nil, false
	}
	return result.val, true
}

func (c *Cache) reapLoop(t time.Duration) {
	ticker := time.NewTicker(t)
	for range ticker.C {
		c.reap(time.Now().UTC(), t)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.data, k)
			fmt.Println("reaped")
		}
	}
}
