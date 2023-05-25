package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	defaultExpirationTime = 24 * time.Hour
	cleanupInterval       = 48 * time.Hour
)

type Cache interface {
	Set(k string, x interface{}, d time.Duration)
	Get(k string) (interface{}, bool)
}

func New() Cache {
	return cache.New(defaultExpirationTime, cleanupInterval)
}
