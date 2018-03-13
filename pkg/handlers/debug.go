package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const DebugPath = "/debug"

func (h *Handlers) Debug(w http.ResponseWriter, r *http.Request) {
	namespaces := h.k8sServer.NamespaceList()

	for _, ns := range namespaces {
		fmt.Fprintf(w, "Namespace: %s \n", ns.Name)

		fmt.Fprintln(w, "\n\nDeployments")
		deployments := h.k8sServer.DeploymentList(ns.Name)
		for _, dep := range deployments {
			fmt.Fprintln(w, dep)
		}

		fmt.Fprintln(w, "\n\nIngresses")
		ingresses := h.k8sServer.IngressList(ns.Name)
		for _, ing := range ingresses {
			fmt.Fprint(w, ing)
			rule := ing.Spec.Rules[0]
			fmt.Fprintf(w, "Host: %s\n", rule.Host)
		}
	}

	fmt.Fprintln(w,"\n\nPodLabels")
	for key, val := range h.config.PodLabels {
		fmt.Fprintf(w, "%s - %s\n", key, val)
	}

	fmt.Fprintln(w,"\n\nEnv")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Fprintf(w, "%s - %s\n", pair[0], pair[1])
	}

}
