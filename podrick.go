package main

import (
	"net/http"

	"github.com/TimWoolford/podrick/pkg/handlers"
	"github.com/TimWoolford/podrick/pkg/server"
	"github.com/TimWoolford/podrick/pkg/config"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.Load("/config/config.yaml")

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))

	handler := handlers.New(*server.New(cfg), *cfg)

	r.HandleFunc(handlers.ReadyPath, handler.Ready)
	r.HandleFunc(handlers.StatusPath, handler.Status)
	r.HandleFunc(handlers.AllNamespacePath, handler.AllNamespaces)
	r.HandleFunc(handlers.NamespacePath, handler.Namespace)
	r.HandleFunc(handlers.DeploymentPath, handler.Deployment)

	http.Handle("/", r)
	http.ListenAndServe(":8082", nil)
}
