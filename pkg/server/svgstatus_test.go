package server

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPodHealth(t *testing.T) {
	podStatus := PodStatus{1, 2}

	health := podStatus.Health()

	assert.Equal(t, "&#9650;&#9661;&#9661;", health)
}

func TestStatusColour(t *testing.T) {
	assert.Equal(t,  "#4c1", Up.colour());
	assert.Equal(t,  "red", Down.colour());
	assert.Equal(t,  "#ffc200", Warn.colour());
}
