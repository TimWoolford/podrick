package main

import (
	"net/http"

	"github.com/TimWoolford/podrick/pkg/handlers"
	"github.com/TimWoolford/podrick/pkg/server"
	"github.com/TimWoolford/podrick/pkg/config"
)

func main() {
	cfg := config.Load()

	fs := http.FileServer(http.Dir("/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	handler := handlers.New(*server.New(cfg))

	http.HandleFunc(handlers.ReadyPath, handler.Ready)
	http.HandleFunc(handlers.NamespacePath, handler.K8s)
	http.HandleFunc(handlers.DeploymentPath, handler.Deployment)

	http.ListenAndServe(":8082", nil)
}
