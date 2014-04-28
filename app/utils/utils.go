package utils

import (
	"io/ioutil"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/yosssi/gogithub"
	"github.com/yosssi/xpress/app/consts"
)

// YamlUnmarshal parses a yaml file and set the result to the out interface.
func YamlUnmarshal(path string, out interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bytes, out); err != nil {
		return err
	}
	return nil
}

// StringSliceHas checks s has target string.
func StringSliceHas(s []string, target string) bool {
	for _, str := range s {
		if str == target {
			return true
		}
	}
	return false
}

// AppendFile appends the file.
func AppendArticleFile(files []string, file string) []string {
	if strings.HasPrefix(file, consts.GitHubArticlesPath+"/") && len(strings.Split(file, "/")) == 2 && !StringSliceHas(files, file) {
		return append(files, file)
	}
	return files
}

// UpdatedFiles returns added, removed or modifed files of the hook.
func UpdatedArticleFiles(hook *gogithub.Hook) []string {
	files := make([]string, 0)
	for _, commit := range hook.Commits {
		for _, file := range commit.Added {
			files = AppendArticleFile(files, file)
		}
		for _, file := range commit.Removed {
			files = AppendArticleFile(files, file)
		}
		for _, file := range commit.Modified {
			files = AppendArticleFile(files, file)
		}
	}
	return files
}
