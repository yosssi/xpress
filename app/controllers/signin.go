package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/yosssi/xpress/app/models"
)

func SigninIndex(w http.ResponseWriter, r *http.Request, app *models.Application) {
	render("./app/views/signin/index.gold", nil, w, r, app)
}

func SigninCallback(w http.ResponseWriter, r *http.Request, app *models.Application) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		handleError(w, r, app, err)
	}
	if err := app.GitHubClient.SetAccessToken(u.Query().Get("code")); err != nil {
		handleError(w, r, app, err)
	}
	app.Logger.Debug(fmt.Sprintf("client: %+v", app.GitHubClient))
	render("./app/views/signin/index.gold", nil, w, r, app)
}
