package cache

import (
	"sync"
	"time"
	"fmt"
)


type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mu sync.Mutex
}


func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
    defer c.mu.Unlock()
    c.entry[key] = cacheEntry{
		val: value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if v, ok := c.entry[key]; ok {
		fmt.Println("Retrieving from cache: ", key)
		return v.val, true
	}
	return nil, false
}

func NewCache(interval time.Duration) *Cache{
	newEntry := make(map[string]cacheEntry)
	new := &Cache{
		entry: newEntry,
	}
	go new.reapLoop(interval)
	return new
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	
	
	for range ticker.C{
		c.mu.Lock()
		
		for key, val := range c.entry{
			if time.Time.Before(val.createdAt, time.Now().Add(-interval)){
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}