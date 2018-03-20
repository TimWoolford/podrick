package output

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAddingOkPodToOkOutcome(t *testing.T) {
	outcome := Outcome{outcomeStatus: OK, podStatus: make(PodStatus)}

	outcome.Add(*NewPod("ok-pod-name", StatusContent{"overallStatus": OK.String()}, nil))

	assert.Equal(t, outcome.outcomeStatus, OK)
	assert.Equal(t, outcome.podStatus[OK], []string{"ok-pod-name"})
}

func TestAddingOkPodToWarnOutcome(t *testing.T) {
	outcome := Outcome{outcomeStatus: WARN, podStatus: make(PodStatus)}

	outcome.Add(*NewPod("ok-pod-name", StatusContent{"overallStatus": OK.String()}, nil))

	assert.Equal(t, outcome.outcomeStatus, WARN)
	assert.Equal(t, outcome.podStatus[OK], []string{"ok-pod-name"})
}

func TestAddingOkPodToFailOutcome(t *testing.T) {
	outcome := Outcome{outcomeStatus: FAIL, podStatus: make(PodStatus)}

	outcome.Add(*NewPod("ok-pod-name", StatusContent{"overallStatus": OK.String()}, nil))

	assert.Equal(t, outcome.outcomeStatus, FAIL)
	assert.Equal(t, outcome.podStatus[OK], []string{"ok-pod-name"})
}

func TestAddingWarnPodToOkOutcome(t *testing.T) {
	outcome := Outcome{outcomeStatus: OK, podStatus: make(PodStatus)}

	outcome.Add(*NewPod("warn-pod-name", StatusContent{"overallStatus": WARN.String()}, nil))

	assert.Equal(t, outcome.outcomeStatus, WARN)
	assert.Equal(t, outcome.podStatus[WARN], []string{"warn-pod-name"})
}

func TestAddingWarnPodToWarnOutcome(t *testing.T) {
	outcome := Outcome{outcomeStatus: WARN, podStatus: make(PodStatus)}

	outcome.Add(*NewPod("warn-pod-name", StatusContent{"overallStatus": WARN.String()}, nil))

	assert.Equal(t, outcome.outcomeStatus, WARN)
	assert.Equal(t, outcome.podStatus[WARN], []string{"warn-pod-name"})
}

func TestAddingWarnPodToFailOutcome(t *testing.T) {
	outcome := Outcome{outcomeStatus: FAIL, podStatus: make(PodStatus)}

	outcome.Add(*NewPod("warn-pod-name", StatusContent{"overallStatus": WARN.String()}, nil))

	assert.Equal(t, outcome.outcomeStatus, FAIL)
	assert.Equal(t, outcome.podStatus[WARN], []string{"warn-pod-name"})
}

func TestAddingFailPodToOkOutcome(t *testing.T) {
	outcome := Outcome{outcomeStatus: OK, podStatus: make(PodStatus)}

	outcome.Add(*NewPod("fail-pod-name", StatusContent{"overallStatus": FAIL.String()}, nil))

	assert.Equal(t, outcome.outcomeStatus, FAIL)
	assert.Equal(t, outcome.podStatus[FAIL], []string{"fail-pod-name"})
}

func TestAddingFailPodToWarnOutcome(t *testing.T) {
	outcome := Outcome{outcomeStatus: WARN, podStatus: make(PodStatus)}

	outcome.Add(*NewPod("fail-pod-name", StatusContent{"overallStatus": FAIL.String()}, nil))

	assert.Equal(t, outcome.outcomeStatus, FAIL)
	assert.Equal(t, outcome.podStatus[FAIL], []string{"fail-pod-name"})
}

func TestAddingFailPodToFailOutcome(t *testing.T) {
	outcome := Outcome{outcomeStatus: FAIL, podStatus: make(PodStatus)}

	outcome.Add(*NewPod("fail-pod-name", StatusContent{"overallStatus": FAIL.String()}, nil))

	assert.Equal(t, outcome.outcomeStatus, FAIL)
	assert.Equal(t, outcome.podStatus[FAIL], []string{"fail-pod-name"})
}
