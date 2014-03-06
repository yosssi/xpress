package utils

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
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
