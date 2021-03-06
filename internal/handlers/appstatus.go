package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TimWoolford/podrick/internal/k8s/endpoints"
	"github.com/TimWoolford/podrick/internal/output"
)

const AppStatusPath = "/status/{namespace}/{name}"

func (h *Handlers) AppStatus(w http.ResponseWriter, r *http.Request) {
	request := Parse(r)

	eps := h.k8sServer.Endpoint(request.Namespace, request.Name)

	out := output.New()
	for _, ep := range eps.ReadyEndpoints(request.Port) {
		out.AddPod(loadPodFrom(ep, request.StatusPath))
	}

	for _, ep := range eps.NotReadyEndpoints(request.Port) {
		out.AddPod(loadPodFrom(ep, request.StatusPath))
	}

	bytes, _ := json.MarshalIndent(&out, "", "  ")

	w.Header()["Content-Type"] = []string{"application/json"}
	w.Write(bytes)
}

func loadPodFrom(ep endpoints.K8sEndpoint, statusPath string) output.Pod {
	resp, err := http.Get(ep.StatusUrl(statusPath))
	if err != nil {
		return *output.FailedPod(ep.Name, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return *output.FailedPod(ep.Name, err)
	}

	statusContent := make(output.StatusContent)
	err = json.Unmarshal(body, &statusContent)

	return *output.NewPod(ep.Name, statusContent, err)
}
