package routes

import (
	"github.com/yosssi/xpress/app/controllers"
	"github.com/yosssi/xpress/app/models"
	"net/http"
)

func Routes(app *models.Application) {
	handleFunc("/", app, controllers.TopIndex)
}

func handleFunc(pattern string, app *models.Application, handler func(w http.ResponseWriter, r *http.Request, app *models.Application)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, app)
	})
}
