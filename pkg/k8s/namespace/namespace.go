package namespace

import "github.com/TimWoolford/podrick/pkg/k8s/deployment"

type K8sNamespace struct {
	name        string
	deployments []deployment.K8sDeployment
}

func New(name string, deployments []deployment.K8sDeployment) *K8sNamespace {
	return &K8sNamespace{name: name, deployments: deployments}
}

func (ns K8sNamespace) Name() string {
	return ns.name
}

func (ns K8sNamespace) Deployments() [] deployment.K8sDeployment {
	return ns.deployments
}
