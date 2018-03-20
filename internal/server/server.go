package server

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/TimWoolford/podrick/internal/config"
)

type K8sServer struct {
	clientSet *kubernetes.Clientset
	config    *config.Config
}

func New(config *config.Config) *K8sServer {
	k8sConfig, err := rest.InClusterConfig()

	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(k8sConfig)

	if err != nil {
		panic(err.Error())
	}

	return &K8sServer{clientSet: clientSet, config: config}
}
