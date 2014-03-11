package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yosssi/xpress/app/models"
	"github.com/yosssi/xpress/app/routes"
)

// main executes main processes.
func main() {
	app, err := models.NewApplication()
	if err != nil {
		log.Fatal(err)
	}

	routes.Routes(app)

	listen(app)
}

// listen starts service listening.
func listen(app *models.Application) {
	app.Logger.Debug(fmt.Sprintf("app.GitHub: %+v", app.GitHub))
	app.Logger.Info(fmt.Sprintf("Listening on port %d.", app.Port()))
	http.ListenAndServe(":"+app.PortString(), nil)
}
