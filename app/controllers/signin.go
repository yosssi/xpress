package controllers

import (
	"fmt"
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
	ghUser, err := app.GitHubClient.GetAuthenticatedUser(accessToken)

	if err != nil {
		handleError(w, r, app, err)
		return
	}

	app.Logger.Debugf("ghUser: %+v", ghUser)

	// get a user.
	result := map[string]interface{}{}
	code, err := app.ElasticsearchClient.Search(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, "", &result)

	if err != nil {
		handleError(w, r, app, err)
		return
	}

	if code != http.StatusOK && code != http.StatusNotFound {
		handleError(w, r, app, fmt.Errorf("Search API's HTTP status code is not OK or NotFound. [code: %d]", code))
		return
	}

	app.Logger.Debugf("result: %+v", result)

	result = map[string]interface{}{}

	userMap := map[string]interface{}{"github_id": ghUser.ID, "access_token": accessToken}
	code, err = app.ElasticsearchClient.Create(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, userMap, &result)
	app.Logger.Debugf("code: %d, result: %+v, err: %+v", code, result, err)

	render("./app/views/signin/index.gold", nil, w, r, app)
}
