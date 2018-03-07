package handlers

import (
	"net/http"
	"fmt"
)

func (h *Handlers) K8s(w http.ResponseWriter, r *http.Request) {
	for _, ns := range h.k8sServer.NamespaceList() {
		fmt.Fprintf(w, "Namespace: %s\n", ns.Name)

		for _, pod := range h.k8sServer.PodList(ns.Name) {
			fmt.Fprintf(w, "Pod: %s\n", pod.Name)
		}
		for _, deployment := range h.k8sServer.DeploymentList(ns.Name) {
			fmt.Fprintf(w, "Deployment: %s\n", deployment.Name())
		}
	}
}