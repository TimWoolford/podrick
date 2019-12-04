package handlers

import (
	"text/template"

	"github.com/TimWoolford/podrick/internal/config"
	"github.com/TimWoolford/podrick/internal/server"
)

type Handlers struct {
	k8sServer server.K8sServer
	template  template.Template
	config    config.Config
}

func New(server server.K8sServer, config config.Config) *Handlers {
	tmpl := template.Must(template.ParseGlob("/template/*"))

	return &Handlers{server, *tmpl, config}
}
