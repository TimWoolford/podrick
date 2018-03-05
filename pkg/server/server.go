package server

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"github.com/TimWoolford/podrick/pkg/deployment"
	"github.com/TimWoolford/podrick/pkg/config"
)

type K8sServer struct {
	clientSet *kubernetes.Clientset
	config *config.Config
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

	return &K8sServer{clientSet: clientSet, config: config }
}

func (s *K8sServer) NamespaceList() ([]v1.Namespace) {
	namespaces, err := s.clientSet.CoreV1().Namespaces().List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return namespaces.Items
}

func (s *K8sServer) PodList(namespace string) ([]v1.Pod) {
	pods, err := s.clientSet.CoreV1().Pods(namespace).List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return pods.Items
}

func (s *K8sServer) Deployment(namespace string, appName string) *deployment.K8sDeployment {
	dep, err := s.clientSet.AppsV1().Deployments(namespace).Get(appName, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	return deployment.New(*dep, *s.config)
}

func (s *K8sServer) DeploymentList(namespace string) ([]deployment.K8sDeployment) {
	namespaces, err := s.clientSet.AppsV1().Deployments(namespace).List(metav1.ListOptions{})

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
