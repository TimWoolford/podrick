package pod

import (
	"k8s.io/api/core/v1"
	"fmt"
)

type K8sPod struct {
	pod v1.Pod
}

func New(pod v1.Pod) *K8sPod {
	return &K8sPod{pod: pod}
}

func (p *K8sPod) Name() string {
	return p.pod.Name
}

func (p *K8sPod) StatusUrl(port int32, statusPath string) string {
	return fmt.Sprintf("http://%s:%d%s", p.pod.Status.PodIP, port, statusPath)
}

func (p *K8sPod) Status() v1.PodPhase {
	return p.pod.Status.Phase
}

func (p *K8sPod) IP() string {
	return p.pod.Status.PodIP
}

func (p *K8sPod) IsRunning() bool {
	return p.pod.Status.Phase == v1.PodRunning
}
