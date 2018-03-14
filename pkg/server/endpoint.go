package server

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/TimWoolford/podrick/pkg/k8s/endpoints"
	"fmt"
)

func (s *K8sServer) Endpoint(namespace string, name string) *endpoints.K8sEndpoints {
	eps, err := s.clientSet.CoreV1().Endpoints(namespace).Get(name, metav1.GetOptions{})

	if err != nil {
		if err.Error() == fmt.Sprintf("endpoints \"%s\" not found", name) {
			return endpoints.Empty()
		}
		panic(err)
	}

	return endpoints.New(*eps)
}
