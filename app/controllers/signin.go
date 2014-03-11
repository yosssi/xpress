package controllers

import (
	"net/http"

	"github.com/yosssi/xpress/app/models"
)

func SigninIndex(w http.ResponseWriter, r *http.Request, app *models.Application) {
	render("./app/views/signin/index.gold", nil, w, r, app)
}

func SigninCallback(w http.ResponseWriter, r *http.Request, app *models.Application) {
	render("./app/views/signin/index.gold", nil, w, r, app)
}
