package controllers

import (
	"fmt"
	"net/http"

	"github.com/yosssi/xpress/app/models"
)

// render renders an HTML.
func render(path string, data *interface{}, w http.ResponseWriter, r *http.Request, app *models.Application) {
	tpl, err := app.Generator.ParseFile(path)
	if err != nil {
		handleError(w, r, app, err)
		return
	}
	if err := tpl.Execute(w, models.NewViewData(app, data)); err != nil {
		handleError(w, r, app, err)
		return
	}
}

// handleError handles an error.
func handleError(w http.ResponseWriter, r *http.Request, app *models.Application, err error) {
	app.Logger.Error(fmt.Sprintf("--- %s %s %s", r.Method, r.URL, err.Error()))
	http.Error(w, app.Msg("errmsg_internal_server_error"), http.StatusInternalServerError)
}
