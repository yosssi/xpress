package controllers

import (
	"github.com/yosssi/xpress/app/models"
	"net/http"
)

func TopIndex(w http.ResponseWriter, r *http.Request, app *models.Application) {
	render("./app/views/top/index.gold", nil, w, r, app)
}

func TopIndex2(w http.ResponseWriter, r *http.Request, app *models.Application) {
	render("./app/views/top/index2.gold", nil, w, r, app)
}
