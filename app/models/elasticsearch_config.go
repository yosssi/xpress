package models

import (
	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/utils"
)

type ElasticsearchConfig struct {
	BaseUrl string `yaml:"baseurl"`
}

// ElasticsearchConfig parses a yaml file, generates a ElasticsearchConfig and returns it.
func NewElasticsearchConfig() (*ElasticsearchConfig, error) {
	config := &ElasticsearchConfig{}
	if err := utils.YamlUnmarshal(consts.ElasticsearchConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
