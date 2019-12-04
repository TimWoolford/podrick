package main

import (
	"net/http"

	"github.com/TimWoolford/podrick/internal/config"
	"github.com/TimWoolford/podrick/internal/handlers"
	"github.com/TimWoolford/podrick/internal/server"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"log"
	"os"
)

func main() {
	cfg := config.Load("/config/config.yaml")
	handler := handlers.New(*server.New(cfg), *cfg)
	log.Println("Podrick is starting")

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))

	r.HandleFunc(handlers.ReadyPath, handler.Ready)
	r.HandleFunc(handlers.StatusPath, handler.Status)
	r.HandleFunc(handlers.AllNamespacePath, handler.AllNamespaces)
	r.HandleFunc(handlers.NamespacePath, handler.Namespace)
	r.HandleFunc(handlers.DeploymentPath, handler.Deployment)
	r.HandleFunc(handlers.AppStatusPath, handler.AppStatus)
	r.HandleFunc(handlers.PodPath, handler.AllPods)

	r.HandleFunc("/{namespace}/http/{name}", handler.AppStatus)

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	http.Handle("/", loggedRouter)
	http.ListenAndServe(":8082", nil)
}
