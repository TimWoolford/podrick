package server

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"

	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"github.com/TimWoolford/podrick/pkg/k8s/deployment"
	"github.com/TimWoolford/podrick/pkg/config"
	"fmt"
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

func (s *K8sServer) NamespaceList() ([]v1.Namespace) {
	namespaces, err := s.clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return namespaces.Items
}

func (s *K8sServer) Deployment(namespace string, name string) *deployment.K8sDeployment {
	dep, err := s.deployment(namespace).Get(name, metav1.GetOptions{})


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
	namespaces, err := s.deployment(namespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	items := namespaces.Items

	vals := make([]deployment.K8sDeployment, len(items))
	for i, v := range items {
		vals[i] = *deployment.New(v, *s.config)
	}
	return vals
}

func (s *K8sServer) deployment(namespace string) appsv1.DeploymentInterface {
	return s.clientSet.AppsV1().Deployments(namespace)
}
