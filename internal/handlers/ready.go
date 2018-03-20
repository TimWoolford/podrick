package handlers

import (
	"fmt"
	"net/http"
)

const ReadyPath = "/ready"

func (h *Handlers) Ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
