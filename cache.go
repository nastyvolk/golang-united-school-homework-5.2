package cache

import "time"

type Cache struct {
	kvPairs map[string]cacheValue
}

type cacheValue struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{kvPairs: make(map[string]cacheValue)}
}

func (c Cache) Get(key string) (string, bool) {
	value, ok := c.kvPairs[key]
	if !ok {
		return "", false
	}
	if value.deadline.IsZero() || value.deadline.After(time.Now()) {
		return value.value, true
	}
	return "", false
}

func (c Cache) Put(key, value string) {
	val := cacheValue{value: value}
	c.kvPairs[key] = val
}

func (c Cache) Keys() []string {
	var keys []string
	for k := range c.kvPairs {
		keys = append(keys, k)
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	val := cacheValue{value: value, deadline: deadline}
	c.kvPairs[key] = val
}
