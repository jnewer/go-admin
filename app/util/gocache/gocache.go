package gocache

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

var (
	cae  *cache.Cache
	once sync.Once
)

func Instance() *cache.Cache {
	once.Do(func() {
		cae = cache.New(1*time.Minute, 3*time.Minute)
	})
	return cae
}
