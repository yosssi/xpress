package jobs

import (
	"fmt"
	"net/http"

	"github.com/yosssi/gogithub"
	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/models"
	"github.com/yosssi/xpress/app/utils"
)

func HookCreate(app *models.Application, hooks <-chan *gogithub.Hook) {
	for hook := range hooks {
		app.Logger.Infof("Hook job starts. [hook: %+v]", hook)

		// Execute only if the commit branch is the master branch.
		if hook.Ref != consts.GitHubRefPrefix+hook.Repository.MasterBranch {
			app.Logger.Infof("Hook process was skipped because the hook's branch was not the master branch.")
			return
		}

		// Get the access token from Elasticsearch.
		searchResult := models.UserSearchResult{}
		code, err := app.ElasticsearchClient.Search(consts.ElasticsearchIndexXpress, consts.ElasticsearchTypeUser, "", &searchResult)
		if err != nil {
			app.Logger.Error(err.Error())
			return
		}

		app.Logger.Debugf("code: %d, searchResult: %+v", code, searchResult)

		if code != http.StatusOK {
			var msg string
			if code == http.StatusNotFound {
				msg = "could not find an user."
			} else {
				msg = fmt.Sprintf("search process ends with an invalid status code. [code: %d]", code)
			}
			app.Logger.Error(msg)
			return
		}

		user := searchResult.User()

		app.Logger.Debugf("user: %+v", user)

		accessToken := user.AccessToken

		if accessToken == "" {
			msg := "could not get an access token."
			app.Logger.Error(msg)
			return
		}

		app.GitHubClient.AccessToken = accessToken

		// Get added, removed or modified files.
		for _, file := range utils.UpdatedArticleFiles(hook) {
			app.Logger.Debugf("file: %s", file)
			content, code, err := app.GitHubClient.GetContent(hook.Repository.Owner.Name, hook.Repository.Name, hook.Repository.MasterBranch, file)
			if err != nil {
				app.Logger.Error(err.Error())
				return
			}
			app.Logger.Debugf("code: %d content:\n%s", code, content)
			switch code {
			case http.StatusNotFound:
			case http.StatusOK:
			}
		}

	}
}
