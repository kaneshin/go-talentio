package talentio

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	conf := NewConfig()
	assert.NotNil(conf)
	assert.Nil(conf.httpClient)
	assert.Empty(conf.accessToken)

	conf.WithHTTPClient(http.DefaultClient).
		WithAccessToken("access-token")

	assert.Equal(http.DefaultClient, conf.httpClient)
	assert.Equal("access-token", conf.accessToken)
}
