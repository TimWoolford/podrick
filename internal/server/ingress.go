package server

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/api/extensions/v1beta1"
)

func (s *K8sServer) IngressList(namespace string) []v1beta1.Ingress {
	ingressList, err := s.clientSet.ExtensionsV1beta1().Ingresses(namespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	return ingressList.Items
}

func (s *K8sServer) Ingress(namespace string, name string) v1beta1.Ingress {
	ingress, err := s.clientSet.ExtensionsV1beta1().Ingresses(namespace).Get(name, metav1.GetOptions{})

	if err != nil {
		panic(err)
	}

	return *ingress
}
