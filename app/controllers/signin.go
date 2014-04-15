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
	searchResult := models.UserSearchResult{}
	// get a user.
	code, err := app.ElasticsearchClient.Search(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, "", &searchResult)
	if err != nil {
		handleError(w, r, app, err)
		return
	}

	app.Logger.Debugf("searchResult: %+v", searchResult)

	var user *models.User = nil

	switch code {
	case http.StatusNotFound:
		createResult := models.CreateResult{}
		userMap := map[string]interface{}{"github_id": ghUser.ID, "access_token": accessToken}
		code, err = app.ElasticsearchClient.Create(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, userMap, &createResult)
		if err != nil {
			handleError(w, r, app, err)
			return
		}
		app.Logger.Debugf("code: %d, createResult: %+v, err: %+v", code, createResult, err)
		user = models.NewUser(createResult.ID, accessToken, ghUser.ID)
	case http.StatusOK:
		user = searchResult.User()
		if ghUser.ID != user.GitHubID {
			app.Logger.Errorf("GitHub account is different from the expected one. [ghUser: %+v][user: %+v]", ghUser, user)
			http.Error(w, app.Msg("errmsg_invalid_account"), http.StatusForbidden)
			return
		}
		if accessToken != user.AccessToken {
			updateResult := models.UpdateResult{}
			code, err = app.ElasticsearchClient.Update(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, user.ID, map[string]interface{}{"doc": map[string]string{"access_token": accessToken}}, &updateResult)
			if err != nil {
				handleError(w, r, app, err)
				return
			}
			if code != http.StatusOK {
				handleError(w, r, app, fmt.Errorf("Update API's HTTP status code is not OK. [code: %d]", code))
				return
			}
			app.Logger.Debugf("code: %d, updateResult: %+v, err: %+v", code, updateResult, err)
			user.AccessToken = accessToken
		}
	default:
		handleError(w, r, app, fmt.Errorf("Search API's HTTP status code is not OK or NotFound. [code: %d]", code))
		return
	}

	app.Logger.Debugf("user: %+v", user)

	render("./app/views/signin/index.gold", nil, w, r, app)
}
