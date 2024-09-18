package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu 		 sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	// Creates a new stuct with the new cachemap and interval time
	c := &Cache{
		cacheMap: make(map[string]cacheEntry),
		mu: sync.Mutex{},
		interval: interval,
	}

	// concurrently running reaploop 
	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	// Locking because we are dealing with a map
	c.mu.Lock()
	defer c.mu.Unlock()

	// Creating a new cache entry with the current time and data
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}

	// Updating the map with that new entry
	c.cacheMap[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// Locking because we are mutating a map
	c.mu.Lock()
	defer c.mu.Unlock()

	// Finding key, if found then return the data and true
	if val, ok := c.cacheMap[key]; ok {
		return val.val, true
	}

	// Not found then return false and no data
	return nil, false
}

func (c *Cache) reapLoop() {
	// Ticker to continuously count every 5 seconds
	ticker := time.NewTicker(5000 * time.Millisecond)
	defer ticker.Stop()

	// Checking each value that is sent by the ticker.C channel
	for range ticker.C {
		currentTime := time.Now()
		c.mu.Lock()

		// Going through the map and comparing to see if there are older entries
		for key, val := range c.cacheMap {
			if currentTime.Sub(val.createdAt) > 5000*time.Millisecond{
				// If entries are older then we delete
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}