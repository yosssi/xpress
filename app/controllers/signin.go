package controllers

import (
	"net/http"
	"net/url"

	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/models"
)

func SigninIndex(w http.ResponseWriter, r *http.Request, app *models.Application) {
	render("./app/views/signin/index.gold", nil, w, r, app)
}

func SigninCallback(w http.ResponseWriter, r *http.Request, app *models.Application) {
	// parse the URL.
	u, err := url.Parse(r.URL.String())

	if err != nil {
		handleError(w, r, app, err)
		return
	}

	// get a GitHub access token.
	accessToken, err := app.GitHubClient.GetAccessToken(u.Query().Get("code"))

	if err != nil {
		handleError(w, r, app, err)
		return
	}

	app.Logger.Debugf("accessToken: %s", accessToken)

	// get a GitHub user.
	user, err := app.GitHubClient.GetAuthenticatedUser(accessToken)

	if err != nil {
		handleError(w, r, app, err)
		return
	}

	app.Logger.Debugf("user: %+v", user)

	// get an account.
	code, account, err := app.ElasticsearchClient.Get(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeAccount, consts.ElasticsearchAccountTopID)

	if err != nil {
		handleError(w, r, app, err)
		return
	}

	app.Logger.Debugf("code: %d, account: %+v", code, account)

	if code == http.StatusOK {

	} else {

	}

	render("./app/views/signin/index.gold", nil, w, r, app)
}
