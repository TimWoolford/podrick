package handlers

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Request struct {
	Namespace  string
	Name       string
	StatusPath string
	Port       int32
}

func Parse(r *http.Request) *Request {
	vars := mux.Vars(r)
	return &Request{
		Namespace:  vars["namespace"],
		Name:       vars["name"],
		Port:       port(r),
		StatusPath: statusPath(r),
	}
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
