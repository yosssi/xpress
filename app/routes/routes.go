package routes

import (
	"fmt"
	"github.com/yosssi/xpress/app/controllers"
	"github.com/yosssi/xpress/app/models"
	"net/http"
	"time"
)

// Routes connects a path to a controller's action.
func Routes(app *models.Application) {
	handleFunc("/", app, controllers.TopIndex)
	handleFunc("/w", app, controllers.TopIndex2)
}

// handleFunc calls a handler.
func handleFunc(pattern string, app *models.Application, handler func(w http.ResponseWriter, r *http.Request, app *models.Application)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info(fmt.Sprintf("--> %s %s", r.Method, r.URL))
		startTime := time.Now()
		handler(w, r, app)
		endTime := time.Now()
		app.Logger.Info(fmt.Sprintf("<-- %s %s %dms", r.Method, r.URL, endTime.Sub(startTime)/time.Millisecond))
	})
}
