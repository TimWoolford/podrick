package server

import (
	"fmt"
	"net/http"
	"text/template"
	"os"
	"strings"
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
		for _, deployment := range f.k8sServer.DeploymentList(ns.Name) {
			fmt.Fprintf(w, "Deployment: %s\n", deployment.ApplicationName())
		}
	}
}

func (f *Handlers) pod(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"image/svg+xml"}

	path := strings.Split(strings.TrimPrefix(r.URL.Path, "/app/"), "/")

	deployment := f.k8sServer.Deployment(path[0], path[1])

	podStatus := PodStatus{deployment.Replicas(), 0}
	status := SvgStatus{
		Version:       deployment.Version(),
		PodHealth:     podStatus.Health(),
		PrimaryColour: "red",
		State:         Up,
	}

	tmpl := template.Must(template.New("output.svg").ParseFiles("pkg/server/web/output.svg"))

	err := tmpl.Execute(w, status)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
