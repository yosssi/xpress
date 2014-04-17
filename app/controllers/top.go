package controllers

import (
	"net/http"

	"github.com/yosssi/xpress/app/models"
)

func TopIndex(w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) bool {
	render("./app/views/top/index.gold", nil, w, r, app, rCtx)
	return false
}
