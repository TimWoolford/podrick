package main

import (
	"github.com/TimWoolford/podrick/pkg/handlers"
	"github.com/TimWoolford/podrick/pkg/server"
	"net/http"
	"github.com/TimWoolford/podrick/pkg/config"
	"log"
)

func main() {
	cfg := config.Load()
	log.Println(cfg)

	handler := handlers.New(*server.New(cfg))

	http.HandleFunc("/ready", handler.Ready)
	http.HandleFunc("/namespaces", handler.K8s)
	http.HandleFunc("/deployment/", handler.Deployment)

	http.ListenAndServe(":8082", nil)
}
