package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewConfig(t *testing.T) {
	expectedConfig := &Config{}
	newConfig := NewConfig()
	assert.Equal(t, expectedConfig, newConfig)
}
