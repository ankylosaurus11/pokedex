package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Cache map[string]cacheEntry
	Mu    *sync.Mutex
}

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Cache[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	entry, ok := c.Cache[key]
	return entry.Val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			<-ticker.C
			c.Reap(interval)
		}
	}()
}

func (c *Cache) Reap(interval time.Duration) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	for key, entry := range c.Cache {
		if time.Since(entry.CreatedAt) > interval {
			delete(c.Cache, key)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		Cache: make(map[string]cacheEntry),
		Mu:    &sync.Mutex{},
	}

	go newCache.reapLoop(interval)

	return newCache
}
