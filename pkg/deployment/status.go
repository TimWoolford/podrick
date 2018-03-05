package deployment

import "strings"

type State int

const upIcon = "&#9650;"
const downIcon = "&#9661;"

const (
	Up   State = iota + 1
	Down
	Warn
)

func (state State) colour() string {
	switch state {
	case Up:
		return "#4c1"
	case Down:
		return "red"
	default:
		return "#ffc200"
	}
}

func (dep *K8sDeployment) PodStatus() (*PodStatus) {
	status := dep.deployment.Status
	return &PodStatus{
		upCount:   int(status.ReadyReplicas),
		downCount: int(status.UnavailableReplicas),
	}
}

type PodStatus struct {
	upCount   int
	downCount int
}

func (status *PodStatus) Health() string {
	return strings.Repeat(upIcon, status.upCount) + strings.Repeat(downIcon, status.downCount)
}

type SvgStatus struct {
	Version       string
	State         State
	PrimaryColour string
	StatusUri     string
	PodHealth     string
}
