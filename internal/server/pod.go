package server

import (
	"context"
	"github.com/TimWoolford/podrick/internal/k8s/pod"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *K8sServer) PodList(namespace string) []pod.K8sPod {
	podList, err := s.clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	pods := make([]pod.K8sPod, len(podList.Items))
	for i, p := range podList.Items {
		pods[i] = *pod.New(p)
	}
	return pods
}

func (s *K8sServer) Pod(namespace string, name string) *pod.K8sPod {
	thePod, err := s.clientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic(err)
	}
	return pod.New(*thePod)
}
