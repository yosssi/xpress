package controllers

import (
	"fmt"
	"net/http"

	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/models"
)

// CommonGetUser gets the user ID from the session and the user from the database.
func CommonGetUser(w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) bool {
	// Get the user's ID from the session.
	store, err := app.NewRediStore()
	if err != nil {
		handleError(w, r, app, fmt.Errorf("An error occurred while calling app.NewRediStore(). [error: %+v]", err))
		return false
	}
	defer store.Close()
	session, err := store.Get(r, app.RediStoreConfig.SessionKey)
	if err != nil {
		if err.Error() == consts.ERR_MSG_SECURECOOKIE_NOT_VALID {
			if err = deleteSession(session, r, w); err != nil {
				handleError(w, r, app, fmt.Errorf("An error occurred while calling deleteSession(). [error: %+v]", err))
				return false
			}
			return true
		} else {
			handleError(w, r, app, fmt.Errorf("An error occurred while calling store.Get(). [error: %+v]", err))
			return false
		}
	}
	userID, prs := session.Values[consts.SessionKeyUserID]
	if !prs {
		return true
	}

	// Get the user.
	userGetResult := models.UserGetResult{}
	code, err := app.ElasticsearchClient.Get(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, userID.(string), &userGetResult)
	if err != nil {
		handleError(w, r, app, fmt.Errorf("An error occurred while calling app.ElasticsearchClient.Get(). [error: %+v]", err))
		return false
	}
	app.Logger.Debugf("code: %+d, userGetResult: %+v", code, userGetResult)
	if code != http.StatusOK {
		if err = deleteSession(session, r, w); err != nil {
			handleError(w, r, app, fmt.Errorf("An error occurred while calling deleteSession(). [error: %+v]", err))
			return false
		}
		return true
	}
	rCtx.User = userGetResult.User()
	app.Logger.Debugf("user: %+v", rCtx.User)

	return true
}

func CommonSignInRequired(w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) bool {
	if rCtx.User == nil {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return false
	}
	return true
}

func CommonNotSignInRequired(w http.ResponseWriter, r *http.Request, app *models.Application, rCtx *models.RequestContext) bool {
	if rCtx.User != nil {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return false
	}
	return true
}
