package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/models"
)

func HookCreate(w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) bool {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, r, app, err)
		return false
	}
	defer r.Body.Close()
	hook := &models.Hook{}
	if err := json.Unmarshal(body, hook); err != nil {
		handleError(w, r, app, err)
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
		handleError(w, r, app, err)
		return false
	}
	app.Logger.Debugf("code %d, searchResult: %+v", code, searchResult)

	fmt.Fprint(w, consts.HookResultProcessed)
	return false
}
