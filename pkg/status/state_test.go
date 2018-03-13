package status

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatusColour(t *testing.T) {
	assert.Equal(t, "#4c1", Up.Colour());
	assert.Equal(t, "red", Fail.Colour());
	assert.Equal(t, "#ffc200", Warn.Colour());
}
