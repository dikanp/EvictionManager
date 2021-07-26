package main

import (
	// "fmt"
	"fmt"
	"memory_cache/models"
)


func main() {
	// cache := models.InMemoryCache{
	// 	Limit: 3,
	// }
	// cache.Add("key1", "value1") 
	// cache.Add("key2", "value2") 
	// cache.Add("key3", "value3") 
	// cache.Add("key2", "value2.1") 
	// cache.Get("key3") 
	// cache.Get("key1") // return value1
	// cache.Get("key3") // return value3
	// cache.Keys() // return ["key1", "key2", "key3"]
	// // cache.Add("key4", "value1") // return / throw Error("key_limit_exceeded")
	// cache.Keys() // return ["key1", "key2", "key3"]
	// cache.Clear() // return 3
	// cache.Keys() // return []
		
	// cache2 := models.InMemoryCache{}
	// cache2.Limit = 3
	// cache2.EvictionManager = &models.LRUEvictionManager{}
	// cache2.Add("key1", "value1") // return 0
	// cache2.Add("key2", "value2") // return 0
	// cache2.Add("key3", "value3") // return 0
	// cache2.Add("key2", "value2.1") // return 1
	// fmt.Println(cache2.Keys())
	// cache2.Get("key3") // return value3
	// cache2.Get("key1") // return value1
	// cache2.Get("key2") // return value2.1
	// cache2.Keys() // return ["key1", "key2", "key3"]
	// cache2.Add("key4", "value4") // return 0
	// fmt.Println(cache2.Keys()) // return ["key1", "key2", "key4"]
	// // (key 3 is the least recently used key, so when key4 added, we will remove key3 from cache)
	// cache2.Clear() // return 3
	// cache2.Keys() // return []
	// // cache.EvictionManager = &models.LFUEvictionManager{}
	// // cache.Add("5", "asd")

	cache3 := models.InMemoryCache{}
	cache3.Limit = 3
	cache3.EvictionManager = &models.LFUEvictionManager{}
	cache3.Add("key1", "value1") // return 0
	cache3.Add("key2", "value2") // return 0
	cache3.Add("key3", "value3") // return 0
	cache3.Add("key2", "value2.1") // return 1
	cache3.Get("key3") // return value3
	cache3.Get("key1") // return value1
	cache3.Get("key2") // return value2.1
	cache3.Get("key3") // return value3
	cache3.Get("key1") // return value1
	// cache3.Get("key2") // return value1
	cache3.Keys() // return ["key1", "key2", "key3"]
	cache3.Add("key4", "value4") // return 0
	cache3.Get("key4") // return value1
	fmt.Println(cache3.Keys()) // return ["key1", "key3", "key4"]
	// (key1 has 2 freq, key 2 has 1 freq, and key 3 has 2 freq, so when key4 added, we will remove key 2 from cache)
	cache3.Clear() // return 3
	cache3.Keys() // return []
}