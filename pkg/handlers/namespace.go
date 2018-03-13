package handlers

import (
	"net/http"

	"github.com/TimWoolford/podrick/pkg/k8s/namespace"
)

const AllNamespacePath = "/namespace"
const NamespacePath = "/namespace/{namespace}"

func (h *Handlers) AllNamespaces(w http.ResponseWriter, r *http.Request) {
	namespaces := h.k8sServer.NamespaceList()

	nss := make([]namespace.K8sNamespace, len(namespaces))

	for i, ns := range namespaces {
		nss[i] = *namespace.New(ns.Name, h.k8sServer.DeploymentList(ns.Name))
	}

	h.template.Lookup("namespaces.html").Execute(w, nss)
}

func (h *Handlers) Namespace(w http.ResponseWriter, r *http.Request) {
	request := Parse(r)

	deployments := h.k8sServer.DeploymentList(request.Namespace)
	ns := namespace.New(request.Namespace, deployments)

	h.template.Lookup("namespace.html").Execute(w, &ns)
}
