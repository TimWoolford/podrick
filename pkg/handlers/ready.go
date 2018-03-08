package handlers

import (
	"net/http"
	"fmt"
)

const ReadyPath = "/ready"

func (h *Handlers) Ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
