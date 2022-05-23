package cache

import "time"

type item struct {
	value    string
	deadline time.Time
}

type Cache struct {
	Items map[string]item
}

func NewCache() Cache {
	return Cache{Items: map[string]item{}}
}

func (i *item) isExpired() bool {
	if i.deadline.IsZero() {
		return false
	}
	now := time.Now()
	return i.deadline.Before(now)
}

func (c Cache) Get(key string) (string, bool) {
	value, ok := c.Items[key]
	if !ok {
		return "", false
	}

	if value.isExpired() {
		delete(c.Items, key)
		return "", false
	}

	return value.value, true
}

func (c *Cache) Put(key, value string) {
	i := item{value: value}
	c.Items[key] = i
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.Items))

	for k := range c.Items {
		_, ok := c.Get(k)
		if ok {
			keys = append(keys, k)
		}
	}

	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	i := item{value: value, deadline: deadline}
	c.Items[key] = i
}
