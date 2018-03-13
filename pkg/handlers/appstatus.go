package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"k8s.io/api/core/v1"

	"github.com/TimWoolford/podrick/pkg/k8s/endpoints"
	"github.com/TimWoolford/podrick/pkg/output"
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
		k8sPod := h.k8sServer.Pod(request.Namespace, ep.Name)
		if k8sPod.Status() == v1.PodRunning {
			out.AddPod(loadPodFrom(ep, request.StatusPath))
		}
	}

	bytes, _ := json.Marshal(&out)

	w.Write(bytes)
}

func loadPodFrom(ep endpoints.K8sEndpoint, statusPath string) output.Pod {
	resp, err := http.Get(ep.StatusUrl(statusPath))
	defer resp.Body.Close()

	if err != nil {
		return *output.FailedPod(ep.Name, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return *output.FailedPod(ep.Name, err)
	}

	statusContent := make(output.StatusContent)
	err = json.Unmarshal(body, &statusContent)

	return *output.NewPod(ep.Name, statusContent, err)
}
