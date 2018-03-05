package deployment

import (
	"k8s.io/api/apps/v1"
)

type K8sDeployment struct {
	deployment v1.Deployment
}

func New(deployment v1.Deployment) *K8sDeployment {
	return &K8sDeployment{deployment: deployment}
}

func (dep *K8sDeployment) Name() (string) {
	return dep.deployment.Name
}

func (dep *K8sDeployment) ApplicationName() (string) {
	return dep.deployment.Labels["app_name"]
}

func (dep *K8sDeployment) Version() (string) {
	appVersion := dep.label("app_version")
	configVersion := dep.label("config_version")

	if len(configVersion) > 0 {
		appVersion += "-" + configVersion
	}
	return appVersion
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
