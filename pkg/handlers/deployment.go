package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const DeploymentPath = "/deployment/{namespace}/{deployment}"
const templateName = "deployment.svg"

func (h *Handlers) Deployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deployment := h.k8sServer.Deployment(vars["namespace"], vars["deployment"])

	w.Header()["Content-Type"] = []string{"image/svg+xml"}
	err := h.template.Lookup(templateName).Execute(w, deployment.SvgStatus())

	if err != nil {
		log.Println(err)
	}
}