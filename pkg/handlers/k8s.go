package handlers

import (
	"net/http"
	"fmt"
)

func (f *Handlers) K8s(w http.ResponseWriter, r *http.Request) {
	for _, ns := range f.k8sServer.NamespaceList() {
		fmt.Fprintf(w, "Namespace: %s\n", ns.Name)

		for _, pod := range f.k8sServer.PodList(ns.Name) {
			fmt.Fprintf(w, "Pod: %s\n", pod.Name)
		}
		for _, deployment := range f.k8sServer.DeploymentList(ns.Name) {
			fmt.Fprintf(w, "Deployment: %s\n", deployment.ApplicationName())
		}
	}
}