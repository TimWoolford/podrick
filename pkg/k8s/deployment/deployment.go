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

func (dep K8sDeployment) Name() string {
	name, present := dep.deployment.Labels[dep.config.AppNameLabel]
	if present {
		return name
	}

	return dep.deployment.Name
}

func (dep K8sDeployment) Namespace() string {
	return dep.deployment.Namespace
}

func (dep K8sDeployment) Version() string {
	versions := make([]string, len(dep.config.VersionLabels))

	i := 0
	for _, label := range dep.config.VersionLabels {
		theLabel, present := dep.deployment.Labels[label]
		if present {
			versions[i] = theLabel
			i = i + 1
		}
	}
	appVersion := strings.Join(versions[0:i], "-")
	if len(appVersion) > 0 {
		return appVersion
	}

	return "Unknown"
}

func (dep K8sDeployment) Status() *status.SvgStatus {
	return &status.SvgStatus{
		ClusterName:   dep.config.ClusterName,
		Version:       dep.Version(),
		PrimaryColour: "green",
		State:         status.Up,
		UpReplicas:    int(dep.deployment.Status.ReadyReplicas),
		DownReplicas:  int(dep.deployment.Status.UnavailableReplicas),
	}
}
