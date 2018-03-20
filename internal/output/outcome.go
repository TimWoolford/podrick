package output

import "encoding/json"

type cleanPodStatus map[string][]string

type Outcome struct {
	outcomeStatus OutcomeStatus
	podStatus     PodStatus
}

func newOutcome() Outcome {
	return Outcome{outcomeStatus: OK, podStatus: make(PodStatus)}
}

func (o Outcome) MarshalJSON() ([]byte, error) {
	podStatus := make(cleanPodStatus)
	for key, val := range o.podStatus {
		podStatus[key.String()] = val
	}

	return json.Marshal(&struct {
		Status    string         `json:"status"`
		PodStatus cleanPodStatus `json:"podStatus"`
	}{
		Status:    o.outcomeStatus.String(),
		PodStatus: podStatus,
	})
}

func (o *Outcome) Add(pod Pod) {
	o.outcomeStatus = o.outcomeStatus.Prioritise(pod.Status())
	o.podStatus[pod.Status()] = append(o.podStatus[pod.Status()], pod.name)
}
