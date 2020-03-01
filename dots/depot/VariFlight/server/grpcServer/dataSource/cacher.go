package VariFlight

import "time"

type cacher struct {
	records map[string]*data
}

func newCacher() *cacher {
	return &cacher{map[string]*data{}}
}

func (c *cacher) create(data *data) {
	if _, ok := c.records[data.token]; !ok {
		c.records[data.token] = data
	}
}

func (c *cacher) read(token string) *data {
	return c.records[token]
}

func (c *cacher) updateUpdateAtTime(token string, newUpdatedAtTime time.Time) {
	if data, ok := c.records[token]; ok {
		data.updatedAtTime = newUpdatedAtTime
	}
}

func (c *cacher) update(token, digest string, updatedAtTime time.Time, value []VariFlightData) {
	if data, ok := c.records[token]; ok {
		data.digest = digest
		data.updatedAtTime = updatedAtTime
		data.value = value
	}
}
