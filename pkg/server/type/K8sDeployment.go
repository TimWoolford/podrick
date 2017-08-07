package _type

import (
	"k8s.io/api/apps/v1beta1"
)

type K8sDeployment struct {
	deployment v1beta1.Deployment
}

func NewDeployment(deployment v1beta1.Deployment) K8sDeployment {
	return K8sDeployment{deployment: deployment}
}

func (dep *K8sDeployment) Name() (string) {
	return dep.deployment.Name
}

func (dep *K8sDeployment) Replicas() (int) {
	return int(*dep.deployment.Spec.Replicas)
}

func (dep *K8sDeployment) ApplicationName() (string) {
	return dep.deployment.Labels["app_name"]
}

func (dep *K8sDeployment) Version() (string) {
	appVersion := dep.label("app_version")
	configVersion := dep.label("config_version")

	if len(configVersion) >0 {
		appVersion += "-" + configVersion
	}
	return appVersion
}

func (dep *K8sDeployment) label(labelName string) (string) {
	return dep.deployment.Labels[labelName]
}
