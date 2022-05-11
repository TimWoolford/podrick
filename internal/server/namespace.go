package server

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/api/core/v1"
)

func (s *K8sServer) NamespaceList() []v1.Namespace {
	namespaces, err := s.clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return namespaces.Items
}
