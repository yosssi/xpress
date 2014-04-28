package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yosssi/gogithub"
	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/models"
	"github.com/yosssi/xpress/app/utils"
)

func HookCreate(w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) bool {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, r, app, err)
		return false
	}
	defer r.Body.Close()
	hook := &gogithub.Hook{}
	if err := json.Unmarshal(body, hook); err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	app.Logger.Debugf("hook: %+v", hook)

	// Execute only if the commit branch is the master branch.
	if hook.Ref != consts.GitHubRefPrefix+hook.Repository.MasterBranch {
		app.Logger.Infof("Hook process was skipped because the hook's branch was not the master branch.")
		fmt.Fprint(w, consts.HookResultSkipped)
		return false
	}

	// Get the access token from Elasticsearch.
	searchResult := models.UserSearchResult{}
	code, err := app.ElasticsearchClient.Search(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, "", &searchResult)
	if err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	app.Logger.Debugf("code: %d, searchResult: %+v", code, searchResult)

	if code != http.StatusOK {
		var msg string
		if code == http.StatusNotFound {
			msg = "could not find an user."
		} else {
			msg = fmt.Sprintf("search process ends with an invalid status code. [code: %d]", code)
		}
		app.Logger.Error(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return false
	}

	user := searchResult.User()

	app.Logger.Debugf("user: %+v", user)

	accessToken := user.AccessToken

	if accessToken == "" {
		msg := "could not get an access token."
		app.Logger.Error(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		return false
	}

	app.GitHubClient.AccessToken = accessToken

	// Get added, removed or modified files.
	for _, file := range utils.UpdatedArticleFiles(hook) {
		app.Logger.Debugf("file: %s", file)
		content, code, err := app.GitHubClient.GetContent(hook.Repository.Owner.Name, hook.Repository.Name, hook.Repository.MasterBranch, file)
		if err != nil {
			app.Logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return false
		}
		app.Logger.Debugf("code: %d content:\n%s", code, content)
		switch code {
		case http.StatusNotFound:
		case http.StatusOK:
		}
	}

	// Send a message to the client.
	fmt.Fprint(w, consts.HookResultProcessed)
	return false
}
