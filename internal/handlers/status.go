package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

const StatusPath = "/status"

type Probe struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type StatusPage struct {
	ApplicationVersion string  `json:"applicationVersion"`
	Hostname           string  `json:"hostname"`
	OverallStatus      string  `json:"overallStatus"`
	Probes             []Probe `json:"probes"`
}

func (h *Handlers) Status(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"text/json"}

	statusPage := StatusPage{
		ApplicationVersion: h.config.PodLabels["app_version"],
		Hostname:           os.Getenv("HOSTNAME"),
		OverallStatus:      "OK",
	}

	bytes, _ := json.Marshal(statusPage)

	w.Write(bytes)
}
