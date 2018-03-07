package status

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatusColour(t *testing.T) {
	assert.Equal(t, "#4c1", Up.colour());
	assert.Equal(t, "red", Down.colour());
	assert.Equal(t, "#ffc200", Warn.colour());
}
