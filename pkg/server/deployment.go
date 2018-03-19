package server

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/TimWoolford/podrick/pkg/k8s/deployment"
	"fmt"
)

func (s *K8sServer) Deployment(namespace string, name string) *deployment.K8sDeployment {
	dep, err := s.clientSet.ExtensionsV1beta1().Deployments(namespace).Get(name, metav1.GetOptions{})

	if err != nil {
		notFoundString := fmt.Sprintf("deployments.apps \"%s\" not found", name)
		if err.Error() == notFoundString {
			return deployment.Empty(*s.config)
		}
		panic(err.Error())
	}

	return deployment.New(*dep, *s.config)

}

func (s *K8sServer) DeploymentList(namespace string) ([]deployment.K8sDeployment) {
	deploymentList, err := s.clientSet.ExtensionsV1beta1().Deployments(namespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	vals := make([]deployment.K8sDeployment, len(deploymentList.Items))
	for i, v := range deploymentList.Items {
		vals[i] = *deployment.New(v, *s.config)
	}
	return vals
}