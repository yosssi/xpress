package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/drone/routes"
	"github.com/yosssi/xpress/app/controllers"
	"github.com/yosssi/xpress/app/models"
)

// Routes connects a path to a controller's action.
func Routes(app *models.Application) {
	mux := routes.New()

	addRoute(routes.GET, "/", mux, app, controllers.TopIndex)
	addRoute(routes.GET, "/signup", mux, app, controllers.SignupIndex)

	pwd, _ := os.Getwd()
	if app.Development() {
		mux.Static("/", pwd)
	} else {
		mux.Static("/public", pwd)
	}

	http.Handle("/", mux)
}

// addRoutes adds a route.
func addRoute(method string, pattern string, mux *routes.RouteMux, app *models.Application, handler func(w http.ResponseWriter, r *http.Request, app *models.Application)) {
	mux.AddRoute(method, pattern, func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info(fmt.Sprintf("--> %s %s", r.Method, r.URL.Path))
		startTime := time.Now()
		handler(w, r, app)
		endTime := time.Now()
		app.Logger.Info(fmt.Sprintf("<-- %s %s %dms", r.Method, r.URL.Path, endTime.Sub(startTime)/time.Millisecond))
	})
}
