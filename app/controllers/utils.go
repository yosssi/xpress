package controllers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/yosssi/xpress/app/models"
)

// render renders an HTML.
func render(path string, data *interface{}, w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) {
	tpl, err := app.Generator.ParseFile(path)
	if err != nil {
		handleError(w, r, app, err)
		return
	}
	if err := tpl.Execute(w, models.NewViewData(app, rCtx, data)); err != nil {
		handleError(w, r, app, err)
		return
	}
}

// handleError handles an error.
func handleError(w http.ResponseWriter, r *http.Request, app *models.Application, err error) {
	app.Logger.Errorf("--- %s %s %s", r.Method, r.URL, err.Error())
	http.Error(w, app.Msg("errmsg_internal_server_error"), http.StatusInternalServerError)
}

// deleteSession deletes the session.
func deleteSession(session *sessions.Session, r *http.Request, w http.ResponseWriter) error {
	maxAge := session.Options.MaxAge
	session.Options.MaxAge = -1
	if err := sessions.Save(r, w); err != nil {
		return err
	}
	session.Options.MaxAge = maxAge
	return nil
}
