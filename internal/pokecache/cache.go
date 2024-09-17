package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	mu 		 sync.Mutex

}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(timeout time.Duration) Cache {
	cacheMap := make(map[string]cacheEntry)
	return Cache{
		cacheMap: cacheMap,
		mu: sync.Mutex{},
	}
}

func (c Cache) Add(key string, val []byte) {
	
}