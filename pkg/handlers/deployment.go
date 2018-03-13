package handlers

import (
	"log"
	"net/http"
)

const DeploymentPath = "/deployment/{namespace}/{name}"
const templateName = "deployment.svg"

func (h *Handlers) Deployment(w http.ResponseWriter, r *http.Request) {
	request := Parse(r)

	deployment := h.k8sServer.Deployment(request.Namespace, request.Name)

	w.Header()["Content-Type"] = []string{"image/svg+xml"}
	err := h.template.Lookup(templateName).Execute(w, deployment.SvgStatus())

	if err != nil {
		log.Println(err)
	}
}
