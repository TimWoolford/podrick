package handlers

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

const DeploymentPath = "/deployment/{namespace}/{deployment}"

func (h *Handlers) Deployment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deployment := h.k8sServer.Deployment(vars["namespace"], vars["deployment"])

	w.Header()["Content-Type"] = []string{"image/svg+xml"}
	err := h.template.Lookup("deployment.svg").Execute(w, deployment.SvgStatus())

	if err != nil {
		log.Println(err)
	}
}