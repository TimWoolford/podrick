package server

import (
	"net/http"
)

func Start() {
	handlers := Handlers{k8sServer: *NewK8sServer()}

	http.HandleFunc("/ready", handlers.ready)
	http.HandleFunc("/namespaces", handlers.k8s)

	http.ListenAndServe(":8082", nil)
}
