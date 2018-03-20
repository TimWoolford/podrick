package status

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPodHealth(t *testing.T) {
	health := SvgStatus{UpReplicas: 1, DownReplicas: 2}.PodHealth()

	assert.Equal(t, "&#9650;&#9661;&#9661;", health)
}
