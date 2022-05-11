package main

import (
	"log"
	"net/http"

	"github.com/TimWoolford/podrick/internal/config"
	"github.com/TimWoolford/podrick/internal/handlers"
	"github.com/TimWoolford/podrick/internal/server"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"os"
)

func main() {
	log.Println("Podrick is starting")
	cfg := config.Load("/config/config.yaml")
	handler := handlers.New(*server.New(cfg), *cfg)

	r := mux.NewRouter()
	r.Use(podrickHeader)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))

	r.HandleFunc(handlers.ReadyPath, handler.Ready)
	r.HandleFunc(handlers.StatusPath, handler.Status)
	r.HandleFunc(handlers.AllNamespacePath, handler.AllNamespaces)
	r.HandleFunc(handlers.NamespacePath, handler.Namespace)
	r.HandleFunc(handlers.DeploymentPath, handler.Deployment)
	r.HandleFunc(handlers.AppStatusPath, handler.AppStatus)
	r.HandleFunc(handlers.PodPath, handler.AllPods)

	r.HandleFunc("/{namespace}/http/{name}", handler.AppStatus)

	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, handlers.AllNamespacePath, http.StatusMovedPermanently)
	})

	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	http.Handle("/", loggedRouter)
	http.ListenAndServe(":8082", nil)
}

func podrickHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header()["X-Source"] = []string{"Podrick"}
		next.ServeHTTP(w, r)
	})
}
