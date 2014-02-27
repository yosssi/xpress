package controllers

import (
	"github.com/yosssi/xpress/app/models"
	"net/http"
)

func TopIndex(w http.ResponseWriter, r *http.Request, app *models.Application) {
	tpl, err := app.Generator.ParseFile("./app/views/top/index.gold")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
