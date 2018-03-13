package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/TimWoolford/podrick/pkg/status"
	"github.com/TimWoolford/podrick/pkg/k8s/endpoints"
)

type PodStatus map[string][]string
type StatusContent map[string]interface{}

type Outcome struct {
	Status    string    `json:"status"`
	PodStatus PodStatus `json:"podStatus"`
}

type Pod struct {
	name       string
	Error      error         `json:"error"`
	StatusPage StatusContent `json:"statusPage"`
}

type Output struct {
	Outcome Outcome        `json:"overallStatus"`
	Pods    map[string]Pod `json:"pods"`
}

func (p Pod) Status() status.OutcomeStatus {
	return status.OutcomeStatusFrom(p.StatusPage["overallStatus"].(string))
}

func newOutput() *Output {
	return &Output{
		Outcome: Outcome{Status: status.OK.String(), PodStatus: make(PodStatus)},
		Pods:    make(map[string]Pod),
	}
}

func (o *Output) AddPod(pod Pod) {
	o.Outcome.Status = status.OutcomeStatusFrom(o.Outcome.Status).Prioritise(pod.Status()).String()
	o.Outcome.PodStatus[pod.Status().String()] = append(o.Outcome.PodStatus[pod.Status().String()], pod.name)
	o.Pods[pod.name] = pod
}

const AppStatusPath = "/status/{namespace}/{deployment}"

func (h *Handlers) AppStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	port := port(r)
	statusPath := statusPath(r)

	eps := h.k8sServer.Endpoint(vars["namespace"], vars["deployment"])

	output := newOutput()
	for _, ep := range eps.ReadyEndpoints(port) {
		pod := loadPodFrom(ep, statusPath)
		output.AddPod(*pod)
	}

	bytes, _ := json.Marshal(&output)
	w.Write(bytes)
}

func loadPodFrom(ep endpoints.K8sEndpoint, statusPath string) *Pod {
	resp, err := http.Get(ep.StatusUrl(statusPath))
	defer resp.Body.Close()
	if err != nil {
		return &Pod{Error: err}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Pod{Error: err}
	}

	statusContent := make(StatusContent)
	err = json.Unmarshal(body, &statusContent)

	return &Pod{name: ep.Name, StatusPage: statusContent, Error: err}
}

func port(r *http.Request) int32 {
	port, err := strconv.ParseInt(r.FormValue("port"), 10, 32)
	if err != nil {
		return 0
	}
	return int32(port)
}

func statusPath(r *http.Request) string {
	path := r.FormValue("statusPath")
	if path == "" {
		return "/status"
	}
	return path
}
