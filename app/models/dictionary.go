package models

import (
	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/utils"
)

type Dictionary map[string]string

// Msg returns a message from the dictionary.
func (d *Dictionary) Msg(s string) string {
	return map[string]string(*d)[s]
}

// NewDictionary parses a yaml file, generates a Dictionary and returns it.
func NewDictionary(locale string) (*Dictionary, error) {
	dictionary := &Dictionary{}
	if err := utils.YamlUnmarshal(consts.DictionariesPath+locale+consts.YmlExtension, dictionary); err != nil {
		return nil, err
	}
	return dictionary, nil
}
