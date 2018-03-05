package deployment

import (
	"k8s.io/api/apps/v1"
	"github.com/TimWoolford/podrick/pkg/config"
	"strings"
)

type K8sDeployment struct {
	deployment v1.Deployment
	config config.Config
}

func New(deployment v1.Deployment, config config.Config) *K8sDeployment {
	return &K8sDeployment{deployment: deployment, config: config}
}

func (dep *K8sDeployment) Name() (string) {
	return dep.deployment.Name
}

func (dep *K8sDeployment) ApplicationName() (string) {
	return dep.deployment.Labels[dep.config.AppNameLabel]
}

func (dep *K8sDeployment) Version() (string) {
	versions := make([]string, len(dep.config.VersionLabels))
	for i, label := range dep.config.VersionLabels {
		versions[i] = dep.label(label)
	}
	return 	strings.Join(versions, "-")
}

func(dep *K8sDeployment) Status() *SvgStatus {
	return &SvgStatus{
		Version:       dep.Version(),
		PodHealth:     dep.PodStatus().Health(),
		PrimaryColour: "green",
		State:         Up,
	}
}

func (dep *K8sDeployment) label(labelName string) (string) {
	return dep.deployment.Labels[labelName]
}
