package main

import (
	"fmt"
	"github.com/yosssi/xpress/app/models"
	"net/http"
)

var (
	application *models.Application
)

// init executes initial processes.
func init() {
	var err error
	application, err = models.NewApplication()
	if err != nil {
		panic(err)
	}
}

// main executes main processes.
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+application.PortString(), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v\n", application)
}
