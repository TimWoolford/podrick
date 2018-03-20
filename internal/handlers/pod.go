package handlers

import (
	"net/http"
	"fmt"
)

const PodPath = "/pod/{namespace}/{name}"

func (h *Handlers) AllPods(w http.ResponseWriter, r *http.Request) {
	request := Parse(r)

	pods := h.k8sServer.PodList(request.Namespace)

	for _, pod := range pods {
		fmt.Fprintf(w, "Pod : %s - %s - %s\n",pod.Name(), pod.Status(), pod.IP())
	}
}
