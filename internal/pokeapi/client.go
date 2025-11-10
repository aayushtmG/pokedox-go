package pokeapi

import (
	"net/http"
	"time"

	"github.com/aayushtmG/pokedexcli/internal"
)

//client
type Client struct {
	httpClient http.Client
	Cache *internal.Cache
}

//new Client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}