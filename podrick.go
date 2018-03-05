package main

import (
	"github.com/TimWoolford/podrick/pkg/handlers"
	"github.com/TimWoolford/podrick/pkg/server"
	"net/http"
)

func main() {
	handler := handlers.New(*server.NewK8sServer())

	http.HandleFunc("/ready", handler.Ready)
	http.HandleFunc("/namespaces", handler.K8s)
	http.HandleFunc("/app/", handler.Pod)

	http.ListenAndServe(":8082", nil)
}
