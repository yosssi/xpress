package main

import (
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
	app.Logger.Debugf("app: %+v", app)
	app.Logger.Debugf("app: %+v", app.Store)
	app.Logger.Debugf("app: %+v", app.Store.Options)
	app.Logger.Debugf("app.GitHubClient: %+v", app.GitHubClient)
	app.Logger.Infof("Listening on port %d.", app.Port())
	http.ListenAndServe(":"+app.PortString(), nil)
}
