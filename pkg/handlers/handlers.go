package handlers

import (
	"text/template"
	"github.com/TimWoolford/podrick/pkg/server"
)

type Handlers struct {
	k8sServer server.K8sServer
	template  template.Template
}

func New(server server.K8sServer) (*Handlers) {
	tmpl := template.Must(template.ParseGlob("/template/*"))

	return &Handlers{server, *tmpl}
}
