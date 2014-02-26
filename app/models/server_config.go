package models

import (
	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/utils"
)

type ServerConfig struct {
	Port        int  `yaml:"port"`
	Development bool `yaml:"development"`
}

// NewServerConfig parses a yaml file, generates a ServerConfig and returns it.
func NewServerConfig() (*ServerConfig, error) {
	config := &ServerConfig{}
	if err := utils.YamlUnmarshal(consts.ServerConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
