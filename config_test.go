package metricmware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigDefault(t *testing.T) {

	cfg := configDefault()
	assert.Nil(t, cfg.Next, "no next")
	assert.Equal(t, "", cfg.Prefix, "no default prefix")

	cfg = configDefault(Config{
		Prefix: "sample",
	})
	assert.Nil(t, cfg.Next, "no next")
	assert.Equal(t, "sample", cfg.Prefix, "sample prefix")

}
