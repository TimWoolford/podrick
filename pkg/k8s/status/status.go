package status

import (
	"strings"
)

const upIcon = "&#9650;"
const downIcon = "&#9661;"

type SvgStatus struct {
	ClusterName   string
	Version       string
	State         State
	PrimaryColour string
	StatusUri     string
	UpReplicas    int
	DownReplicas  int
}

func (status SvgStatus) PodHealth() string {
	return strings.Repeat(upIcon, int(status.UpReplicas)) + strings.Repeat(downIcon, int(status.DownReplicas))
}
