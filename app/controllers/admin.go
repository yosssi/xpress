package controllers

import (
	"net/http"

	"github.com/yosssi/xpress/app/models"
)

func AdminIndex(w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) bool {
	render("./app/views/admin/index.gold", nil, w, r, app, rCtx)
	return false
}
