package output

type PodStatus map[OutcomeStatus][]string
type PodList map[string]Pod

type Output struct {
	Outcome Outcome `json:"overallStatus"`
	Pods    PodList `json:"pods"`
}

func New() *Output {
	return &Output{
		Outcome: newOutcome(),
		Pods:    make(PodList),
	}
}

func (o *Output) AddPod(pod Pod) {
	o.Outcome.Add(pod)
	o.Pods[pod.name] = pod
}
