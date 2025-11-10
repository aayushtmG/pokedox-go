package pokeapi

import (
	"net/http"
	"time"

	"github.com/aayushtmG/pokedexcli/internal/pokecache"
)

//client
type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

//new Client
func NewClient(timeout,cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}