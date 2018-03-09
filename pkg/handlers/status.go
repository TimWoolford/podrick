package handlers

import (
	"net/http"
	"encoding/json"
)

const StatusPath = "/status"

type StatusPage struct {
	ApplicationVersion string `json:"applicationVersion"`
	PropertiesVersion  string `json:"propertiesVersion"`
	OverallStatus      string `json:"overallStatus"`
}

func (h *Handlers) Status(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"text/json"}

	statusPage := StatusPage{
		ApplicationVersion: h.config.PodLabels["app_version"],
		PropertiesVersion:  h.config.PodLabels["config_version"],
		OverallStatus:      "OK",
	}

	bytes, _ := json.Marshal(statusPage)

	w.Write(bytes)
}
