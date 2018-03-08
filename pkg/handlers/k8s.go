package handlers

import (
	"net/http"
	"github.com/TimWoolford/podrick/pkg/k8s/namespace"
)

const NamespacePath = "/namespaces"

func (h *Handlers) K8s(w http.ResponseWriter, r *http.Request) {
	namespaces := h.k8sServer.NamespaceList()

	nss := make([]namespace.K8sNamespace, len(namespaces))

	for i, ns := range namespaces {
		nss[i] = *namespace.New(ns.Name, h.k8sServer.DeploymentList(ns.Name))
	}

	h.template.Lookup("namespaces.html").Execute(w, nss)
}