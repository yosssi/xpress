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
	accessToken, err := app.GitHubClient.GetAccessToken(u.Query().Get("code"))
	if err != nil {
		handleError(w, r, app, err)
	}
	app.Logger.Debug(fmt.Sprintf("accessToken: %s", accessToken))
	user, err := app.GitHubClient.GetAuthenticatedUser(accessToken)
	if err != nil {
		handleError(w, r, app, err)
	}
	app.Logger.Debug(fmt.Sprintf("user: %+v", user))
	render("./app/views/signin/index.gold", nil, w, r, app)
}
