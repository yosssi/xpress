package routes

import (
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
	addRoute(routes.GET, "/admin", mux, app, controllers.CommonGetUser, controllers.CommonSignInRequired, controllers.AdminIndex)
	addRoute(routes.POST, "/hook", mux, app, controllers.HookCreate)
	addRoute(routes.GET, "/signin", mux, app, controllers.CommonGetUser, controllers.CommonNotSignInRequired, controllers.SigninIndex)
	addRoute(routes.GET, "/signin/callback", mux, app, controllers.CommonGetUser, controllers.CommonNotSignInRequired, controllers.SigninCallback)

	pwd, _ := os.Getwd()
	if app.Development() {
		mux.Static("/", pwd)
	} else {
		mux.Static("/static", pwd)
	}

	http.Handle("/", mux)
}

// addRoutes adds a route.
func addRoute(method string, pattern string, mux *routes.RouteMux, app *models.Application, handlers ...func(http.ResponseWriter, *http.Request, *models.Application, *models.RequestContext) bool) {
	mux.AddRoute(method, pattern, func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		app.Logger.Infof("--> %s %s", method, path)
		rCtx := models.NewRequestContext()
		startTime := time.Now()
		for _, handler := range handlers {
			if !handler(w, r, app, rCtx) {
				break
			}
		}
		endTime := time.Now()
		app.Logger.Infof("<-- %s %s %dms", method, path, endTime.Sub(startTime)/time.Millisecond)
	})
}
