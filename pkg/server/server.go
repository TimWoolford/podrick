package server

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type K8sServer struct {
	clientSet *kubernetes.Clientset
}

func NewK8sServer() *K8sServer {
	config, err := rest.InClusterConfig()

	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	return &K8sServer{ clientSet: clientSet }
}

func (s *K8sServer) NamespaceList() ([]v1.Namespace) {
	namespaces, err := s.clientSet.Namespaces().List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return namespaces.Items
}

func (s *K8sServer) PodList(namespace string) ([]v1.Pod) {
	pods, err := s.clientSet.Pods(namespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return pods.Items
}
