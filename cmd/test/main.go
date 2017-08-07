package main

import (
	"net/http"
	"text/template"

	"github.com/TimWoolford/podrick/pkg/server"
	"fmt"
	"os"
)

func main() {
	http.HandleFunc("/pod", pod)
	http.ListenAndServe(":8082", nil)
}

func pod(w http.ResponseWriter, r *http.Request) {

	w.Header()["Content-Type"] = []string{"image/svg+xml"}
	podStatus := server.PodStatus{3, 0}
	status := server.SvgStatus{
		Version:       "111-2",
		PodHealth:     podStatus.Health(),
		PrimaryColour: "red",
		State:         server.Up,
	}

	tmpl := template.Must(template.New("output.svg").ParseFiles("pkg/server/web/output.svg"))

	err := tmpl.Execute(w, status)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	fmt.Fprintln(os.Stdout, tmpl.Name())
}
