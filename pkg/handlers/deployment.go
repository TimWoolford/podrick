package handlers

import (
	"net/http"
	"strings"
	"log"
)

func (h *Handlers) Deployment(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"image/svg+xml"}

	path := strings.Split(strings.TrimPrefix(r.URL.Path, "/deployment/"), "/")

	deployment := h.k8sServer.Deployment(path[0], path[1])

	err := h.template.Execute(w, deployment.Status())

	if err != nil {
		log.Println(err)
	}
}