package output

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	assert.Equal(t, OK, OutcomeStatusFrom("OK"))
	assert.Equal(t, WARN, OutcomeStatusFrom("WARN"))
	assert.Equal(t, FAIL, OutcomeStatusFrom("FAIL"))
	assert.Panics(t, func() { OutcomeStatusFrom("FOO") })
}

func TestOutcomeName(t *testing.T) {
	assert.Equal(t, "OK", OK.String())
	assert.Equal(t, "WARN", WARN.String())
	assert.Equal(t, "FAIL", FAIL.String())
}

func TestOutcomePrioritise(t *testing.T) {
	assert.Equal(t, OK, OK.Prioritise(OK))
	assert.Equal(t, WARN, OK.Prioritise(WARN))
	assert.Equal(t, FAIL, OK.Prioritise(FAIL))

	assert.Equal(t, WARN, WARN.Prioritise(OK))
	assert.Equal(t, WARN, WARN.Prioritise(WARN))
	assert.Equal(t, FAIL, WARN.Prioritise(FAIL))

	assert.Equal(t, FAIL, FAIL.Prioritise(OK))
	assert.Equal(t, FAIL, FAIL.Prioritise(WARN))
	assert.Equal(t, FAIL, FAIL.Prioritise(FAIL))
}
