package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yosssi/gogithub"
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
	hook := &gogithub.Hook{}
	if err := json.Unmarshal(body, hook); err != nil {
		app.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	app.HookC <- hook
	fmt.Fprintf(w, consts.HookResultProcessed)
	return true
}
