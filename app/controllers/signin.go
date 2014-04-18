package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/sessions"
	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/models"
)

func SigninIndex(w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) bool {
	render("./app/views/signin/index.gold", nil, w, r, app, rCtx)
	return false
}

func SigninCallback(w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) bool {
	// parse the URL.
	u, err := url.Parse(r.URL.String())
	if err != nil {
		handleError(w, r, app, err)
		return false
	}
	paramCode := u.Query().Get("code")
	if paramCode == "" {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return false
	}
	// get a GitHub access token.
	err = app.GitHubClient.SetAccessToken(paramCode)
	if err != nil {
		handleError(w, r, app, err)
		return false
	}
	accessToken := app.GitHubClient.AccessToken
	app.Logger.Debugf("accessToken: %s", accessToken)
	// get a GitHub user.
	ghUser, err := app.GitHubClient.GetAuthenticatedUser()
	if err != nil {
		handleError(w, r, app, err)
		return false
	}
	app.Logger.Debugf("ghUser: %+v", ghUser)
	searchResult := models.UserSearchResult{}
	// get a user.
	code, err := app.ElasticsearchClient.Search(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, "", &searchResult)
	if err != nil {
		handleError(w, r, app, err)
		return false
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
			return false
		}
		app.Logger.Debugf("code: %d, createResult: %+v", code, createResult)
		user = models.NewUser(createResult.ID, accessToken, ghUser.ID)
	case http.StatusOK:
		user = searchResult.User()
		if ghUser.ID != user.GitHubID {
			app.Logger.Errorf("GitHub account is different from the expected one. [ghUser: %+v][user: %+v]", ghUser, user)
			http.Error(w, app.Msg("errmsg_invalid_account"), http.StatusForbidden)
			return false
		}
		if accessToken != user.AccessToken {
			updateResult := models.UpdateResult{}
			code, err = app.ElasticsearchClient.Update(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, user.ID, map[string]interface{}{"doc": map[string]string{"access_token": accessToken}}, &updateResult)
			if err != nil {
				handleError(w, r, app, err)
				return false
			}
			if code != http.StatusOK {
				handleError(w, r, app, fmt.Errorf("Update API's HTTP status code is not OK. [code: %d]", code))
				return false
			}
			app.Logger.Debugf("code: %d, updateResult: %+v, err: %+v", code, updateResult, err)
			user.AccessToken = accessToken
		}
	default:
		handleError(w, r, app, fmt.Errorf("Search API's HTTP status code is not OK or NotFound. [code: %d]", code))
		return false
	}

	app.Logger.Debugf("user: %+v", user)

	// Set the user's ID to the session.
	store, err := app.NewRediStore()
	if err != nil {
		handleError(w, r, app, fmt.Errorf("An error occurred while calling app.NewRediStore(). [error: %+v]", err))
		return false
	}
	defer store.Close()
	session, err := store.Get(r, app.RediStoreConfig.SessionKey)
	if err != nil {
		if err.Error() == consts.ErrMsgSecurecookieNotValid {
			if err = deleteSession(session, r, w); err != nil {
				handleError(w, r, app, fmt.Errorf("An error occurred while calling deleteSession(). [error: %+v]", err))
				return false
			}
		} else {
			handleError(w, r, app, fmt.Errorf("An error occurred while calling store.Get(). [error: %+v]", err))
			return false
		}
	}
	session.Values[consts.SessionKeyUserID] = user.ID
	if err = sessions.Save(r, w); err != nil {
		handleError(w, r, app, fmt.Errorf("An error occurred while calling sessions.Save(). [error: %+v]", err))
		return false
	}

	http.Redirect(w, r, "/admin", http.StatusFound)
	return false
}
