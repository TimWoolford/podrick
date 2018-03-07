package deployment

import (
	"k8s.io/api/apps/v1"
	"github.com/TimWoolford/podrick/pkg/config"
	"github.com/TimWoolford/podrick/pkg/k8s/status"
	"strings"
)

type K8sDeployment struct {
	deployment v1.Deployment
	config     config.Config
}

func New(deployment v1.Deployment, config config.Config) *K8sDeployment {
	return &K8sDeployment{deployment: deployment, config: config}
}

func (dep K8sDeployment) Name() (string) {
	name, present := dep.deployment.Labels[dep.config.AppNameLabel]
	if present {
		return name
	}

	return dep.deployment.Name
}

func (dep K8sDeployment) Version() (string) {
	versions := make([]string, len(dep.config.VersionLabels))
	for i, label := range dep.config.VersionLabels {
		versions[i] = dep.deployment.Labels[label]
	}
	return strings.Join(versions, "-")
}

func (dep K8sDeployment) Status() *status.SvgStatus {
	return &status.SvgStatus{
		Version:       dep.Version(),
		PrimaryColour: "green",
		State:         status.Up,
		UpReplicas:    int(dep.deployment.Status.ReadyReplicas),
		DownReplicas:  int(dep.deployment.Status.UnavailableReplicas),
	}
}
