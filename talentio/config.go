package talentio

import "net/http"

type (
	// A Config provides service configuration for service.
	Config struct {
		httpClient  *http.Client
		accessToken string
	}
)

var (
	defaultConfig = *(NewConfig().WithHTTPClient(http.DefaultClient))
)

// NewConfig returns a pointer of new Config objects.
func NewConfig() *Config {
	return &Config{}
}

// WithHTTPClient sets a config HTTPClient value returning a Config pointer
// for chaining.
func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.httpClient = client
	return c
}

// WithAccessToken sets a access token value to verify service returning
// a Config pointer for chaining.
func (c *Config) WithAccessToken(token string) *Config {
	c.accessToken = token
	return c
}
