package handlers

import (
	"net/http"
	"fmt"
)

func (f *Handlers) Ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
