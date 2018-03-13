package deployment

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsV1 "k8s.io/api/apps/v1"

	"github.com/TimWoolford/podrick/pkg/config"
	"github.com/TimWoolford/podrick/pkg/status"
	"strings"
)

type K8sDeployment struct {
	deployment appsV1.Deployment
	config     config.Config
}

func New(deployment appsV1.Deployment, config config.Config) *K8sDeployment {
	return &K8sDeployment{deployment: deployment, config: config}
}

func Empty(config config.Config) *K8sDeployment {
	return &K8sDeployment{
		deployment: appsV1.Deployment{
			ObjectMeta: metaV1.ObjectMeta{Name: "name", Labels: map[string]string{"version": "DOWN"}},
		},
		config: config,
	}
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

func (dep *K8sDeployment) State() status.State {
	switch dep.deployment.Labels["version"] {
	case "DOWN":
		return status.Down
	default:
		return status.Up
	}
}

func (dep K8sDeployment) SvgStatus() *status.SvgStatus {
	return &status.SvgStatus{
		ClusterName:   dep.config.ClusterName,
		Version:       dep.Version(),
		PrimaryColour: dep.State().Colour(),
		UpReplicas:    int(dep.deployment.Status.ReadyReplicas),
		DownReplicas:  int(dep.deployment.Status.UnavailableReplicas),
	}
}
