package pokeapi

import (
	"net/http"
	"time"
	"github.com/montruh-afk/pokedex/internal"
)

// Client -
type Client struct {
	cash *cache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cash: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
