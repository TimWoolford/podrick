package output

import (
	"encoding/json"
)

type PodStatus map[OutcomeStatus][]string
type PodList map[string]Pod

type Outcome struct {
	Status    OutcomeStatus
	PodStatus PodStatus
}

func (o Outcome) MarshalJSON() ([]byte, error) {
	podStatus := make(map[string][]string)
	for key, val := range o.PodStatus {
		podStatus[key.String()] = val
	}
	return json.Marshal(&struct {
		Status    string              `json:"status"`
		PodStatus map[string][]string `json:"podStatus"`
	}{
		Status:    o.Status.String(),
		PodStatus: podStatus,
	})
}

type Output struct {
	Outcome Outcome `json:"overallStatus"`
	Pods    PodList `json:"pods"`
}

func New() *Output {
	return &Output{
		Outcome: Outcome{Status: OK, PodStatus: make(PodStatus)},
		Pods:    make(PodList),
	}
}

func (o *Output) AddPod(pod Pod) {
	o.Outcome.Status = o.Outcome.Status.Prioritise(pod.Status())
	o.Outcome.PodStatus[pod.Status()] = append(o.Outcome.PodStatus[pod.Status()], pod.name)
	o.Pods[pod.name] = pod
}
