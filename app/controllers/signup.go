package controllers

import (
	"github.com/yosssi/xpress/app/models"
	"net/http"
)

func SignupIndex(w http.ResponseWriter, r *http.Request, app *models.Application) {
	render("./app/views/signup/index.gold", nil, w, r, app)
}
