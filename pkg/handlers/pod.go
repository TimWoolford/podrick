package handlers

import (
	"net/http"
	"strings"
	"fmt"
	"os"
)

func (f *Handlers) Pod(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"image/svg+xml"}

	path := strings.Split(strings.TrimPrefix(r.URL.Path, "/app/"), "/")

	deployment := f.k8sServer.Deployment(path[0], path[1])

	err := f.template.Execute(w, deployment.Status())

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}