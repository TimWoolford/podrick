package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLocationName(t *testing.T) {

	config := Load("./test-data/good-test-config.yaml")

	assert.Equal(t, "HERE", config.ClusterName)
}

func TestUnknownLocationName(t *testing.T) {

	config := Load("./test-data/naughty-test-config.yaml")

	assert.Equal(t, "UNK", config.ClusterName)
}
