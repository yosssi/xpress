package controllers

import (
	"net/http"

	"github.com/yosssi/xpress/app/models"
)

func SignupIndex(w http.ResponseWriter, r *http.Request, app *models.Application) {
	render("./app/views/signup/index.gold", nil, w, r, app)
}
