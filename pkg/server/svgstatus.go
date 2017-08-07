package server

import (
	"strings"
)

type State int

const upIcon string = "&#9650;"
const downIcon string = "&#9661;"

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

type PodStatus struct {
	UpCount   int
	DownCount int
}

func (status *PodStatus) Health() string {
	return strings.Repeat(upIcon, status.UpCount) + strings.Repeat(downIcon, status.DownCount)
}

type SvgStatus struct {
	Version       string
	State         State
	PrimaryColour string
	StatusUri     string
	PodHealth     string
}
