package models

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	MockInMemoryCacheNoneEviction	InMemoryCache = InMemoryCache{Limit: 3}
	MockInMemoryCacheLRUEviction	InMemoryCache = InMemoryCache{Limit: 3, EvictionManager: &LRUEvictionManager{}}
	MockInMemoryCacheLFUEviction	InMemoryCache = InMemoryCache{Limit: 3, EvictionManager: &LFUEvictionManager{}}
)

func TestAdd(t *testing.T) {	
	if MockInMemoryCacheNoneEviction.Add("key1", "value1") != 0 {
		t.Errorf("SALAH! harusnya %d", 0)
	}
	
	if (len(MockInMemoryCacheNoneEviction.data)) != 1 {
		t.Errorf("SALAH! harusnya %d", 1)
	}
	
	if MockInMemoryCacheNoneEviction.Add("key1", "value1") != 1 {
		t.Errorf("SALAH! harusnya %d", 1)
	}
	
	if (len(MockInMemoryCacheNoneEviction.data)) != 1 {
		t.Errorf("SALAH! harusnya %d", 1)
	}
}

func TestGet(t *testing.T) {
	MockInMemoryCacheNoneEviction.Add("key1", "value1")
	MockInMemoryCacheNoneEviction.Add("key2", "value1")
	MockInMemoryCacheNoneEviction.Add("key3", "value1")
	MockInMemoryCacheNoneEviction.Get("key2")
}
func TestClear(t *testing.T) {
	t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key1", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key2", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key3", "value1"))
	assert.Equal(t, MockInMemoryCacheLRUEviction.Clear(), 3)
}

func TestKeys(t *testing.T) {
	t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key1", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key2", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key3", "value1"))
	assert.Contains(t, MockInMemoryCacheLRUEviction.Keys(), "key3")
	assert.Contains(t, MockInMemoryCacheLRUEviction.Keys(), "key1")
	assert.Contains(t, MockInMemoryCacheLRUEviction.Keys(), "key2")
	assert.Equal(t, len(MockInMemoryCacheLRUEviction.Keys()), 3)
}

func TestAddKeyLimitExceeded(t *testing.T) {
	
	t.Logf("Add data %2d", MockInMemoryCacheNoneEviction.Add("key1", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheNoneEviction.Add("key2", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheNoneEviction.Add("key3", "value1"))

	// if MockInMemoryCacheNoneEviction.Add("key4", "value1") != -1 {
		// t.Errorf("SALAH! harusnya %d", -1)
		// }
	}
	
	func TestAddLRUNoData(t *testing.T) {
		
		t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key1", "value1"))
		t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key2", "value1"))
		t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key3", "value1"))
		
		if MockInMemoryCacheLRUEviction.Add("key4", "value1") != -1 {
			t.Errorf("SALAH! harusnya %d", -1)
		}
	}
	
	func TestGetLRU(t *testing.T) {
		t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key1", "value1"))
		t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key2", "value1"))
		t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key3", "value1"))
		if MockInMemoryCacheLRUEviction.Get("key1") != "value1" {
			t.Errorf("SALAH! harusnya %d", -1)
	}
}

func TestAddPushPopLRU(t *testing.T) {
	t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key1", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key2", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLRUEviction.Add("key3", "value1"))
	MockInMemoryCacheLRUEviction.Get("key1")
	MockInMemoryCacheLRUEviction.Get("key2")
	MockInMemoryCacheLRUEviction.Get("key3")
	MockInMemoryCacheLRUEviction.Add("key4", "value1")
	// Delete key1
	assert.NotContains(t, MockInMemoryCacheLRUEviction.data, "key1")
	assert.Equal(t, len(MockInMemoryCacheLRUEviction.data), 3)
}

func TestAddLFUNoData(t *testing.T) {
	
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key1", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key2", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key3", "value1"))
	
	if MockInMemoryCacheLFUEviction.Add("key4", "value1") != -1 {
		t.Errorf("SALAH! harusnya %d", -1)
	}
}
func TestGetLFU(t *testing.T) {
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key1", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key2", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key3", "value1"))
	assert.Equal(t, MockInMemoryCacheLFUEviction.Get("key1"), "value1")
}

func TestAddPushPopLFU(t *testing.T) {
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key1", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key2", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key3", "value1"))
	MockInMemoryCacheLFUEviction.Get("key1")
	MockInMemoryCacheLFUEviction.Get("key2")
	MockInMemoryCacheLFUEviction.Get("key3")
	MockInMemoryCacheLFUEviction.Get("key1")
	MockInMemoryCacheLFUEviction.Get("key1")
	MockInMemoryCacheLFUEviction.Get("key2")
	MockInMemoryCacheLFUEviction.Get("key2")
	MockInMemoryCacheLFUEviction.Get("key2")
	MockInMemoryCacheLFUEviction.Get("key3")
	MockInMemoryCacheLFUEviction.Add("key4", "value1")
	// Delete key1
	assert.NotContains(t, MockInMemoryCacheLFUEviction.data, "key3")
	assert.Equal(t, len(MockInMemoryCacheLFUEviction.data), 3)
}

func TestClearLFU(t *testing.T) {
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key1", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key2", "value1"))
	t.Logf("Add data %2d", MockInMemoryCacheLFUEviction.Add("key3", "value1"))
	MockInMemoryCacheLFUEviction.Get("key1")
	MockInMemoryCacheLFUEviction.Get("key2")
	MockInMemoryCacheLFUEviction.Get("key3")
	MockInMemoryCacheLFUEviction.Add("key4", "value1")
	// Delete key1
	MockInMemoryCacheLFUEviction.Clear()
	assert.Equal(t, len(MockInMemoryCacheLFUEviction.data), 0)
}