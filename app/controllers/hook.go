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
		fmt.Fprint(w, "")
		return false
	}

	fmt.Fprint(w, "")
	return false
}
