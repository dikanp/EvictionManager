package models

import (
	"errors"
	"log"
)


type InMemoryCache struct {
	Limit           int
	EvictionManager EvictionManager
	data            map[string]string
}

// func (c *InMemoryCache) Init(limit int, eviction EvictionManager) int {
// 	c.limit = limit
// 	c.evictionManager = eviction
// 	return 0
// }

func (cache *InMemoryCache) Add(key string, value string) int {

	if cache.data[key] != "" {
		cache.data[key] = value
		return 1
	}
	
	if (cache.Limit == len(cache.data)) && (len(cache.data) > 0) {
		if cache.EvictionManager == nil {
			log.Fatal(errors.New("key_limit_exceeded"))
		} else {
			removedData := cache.EvictionManager.pop()
			if removedData == "no data" {
				return -1
			}
			delete(cache.data, removedData)
		}
	}

	if cache.data == nil {
		cache.data = make(map[string]string)
	}
	cache.data[key] = value
	return 0
}

func (cache *InMemoryCache) Get(key string) string {
	if (cache.EvictionManager != nil) {
		cache.EvictionManager.push(key)
	}
	return cache.data[key]
}

func (cache *InMemoryCache) Clear() int{
	countKeysDeleted := len(cache.data)
	if (cache.EvictionManager != nil) {
		cache.EvictionManager.clear()
	}
	cache.data = make(map[string]string)
	return countKeysDeleted
}

func (cache *InMemoryCache) Keys() []string{
	var arrayKeys []string
	for key := range cache.data {
		arrayKeys = append(arrayKeys, key)
	}
	return arrayKeys
}
