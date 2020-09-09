package server_test

import (
	"testing"

	"github.com/openmind13/link-shortener/app/server"
	"github.com/stretchr/testify/assert"
)

func Test_NewConfig(t *testing.T) {
	expectedConfig := &server.Config{}
	newConfig := server.NewConfig()
	assert.Equal(t, expectedConfig, newConfig)
}
