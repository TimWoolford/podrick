package server

import (
	"context"
	"fmt"
	"github.com/TimWoolford/podrick/internal/k8s/endpoints"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *K8sServer) Endpoint(namespace string, name string) *endpoints.K8sEndpoints {
	eps, err := s.clientSet.CoreV1().Endpoints(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		if err.Error() == fmt.Sprintf("endpoints \"%s\" not found", name) {
			return endpoints.Empty()
		}
		panic(err)
	}

	return endpoints.New(eps, s.config)
}
