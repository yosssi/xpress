package utils

import (
	"io/ioutil"
	"launchpad.net/goyaml"
)

// YamlUnmarshal parses a yaml file and set the result to the out interface.
func YamlUnmarshal(path string, out interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := goyaml.Unmarshal(bytes, out); err != nil {
		return err
	}
	return nil
}
