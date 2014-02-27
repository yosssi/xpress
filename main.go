package main

import (
	"fmt"
	"github.com/yosssi/xpress/app/models"
	"github.com/yosssi/xpress/app/routes"
	"net/http"
)

// main executes main processes.
func main() {
	app, err := models.NewApplication()
	if err != nil {
		panic(err)
	}

	routes.Routes(app)

	listen(app)
}

// listen starts service listening.
func listen(app *models.Application) {
	app.Logger.Info(fmt.Sprintf("Listening on port %d.", app.Port()))
	http.ListenAndServe(":"+app.PortString(), nil)
}
