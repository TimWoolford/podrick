package server

import (
	"fmt"
	"net/http"
)

type Handlers struct {
	k8sServer K8sServer
}

func (f *Handlers) ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func (f *Handlers) k8s(w http.ResponseWriter, r *http.Request) {
	for _, ns := range f.k8sServer.NamespaceList() {
		fmt.Fprintf(w, "Namespace: %s\n", ns.Name)

		for _, pod := range f.k8sServer.PodList(ns.Name) {
			fmt.Fprintf(w, "Pod: %s\n", pod.Name)
		}
	}
}
