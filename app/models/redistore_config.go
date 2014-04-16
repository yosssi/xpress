package models

import (
	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/utils"
)

// A RedistoreConfig represents a Redistore config.
type RedistoreConfig struct {
	Size     int    `yaml:"size"`
	Network  string `yaml:"network"`
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	MaxAge   int    `yaml:"max_age"`
}

// NewRedistoreConfig parses a yaml file, generates a RedistoreConfig and returns it.
func NewRedistoreConfig() (*RedistoreConfig, error) {
	config := &RedistoreConfig{}
	if err := utils.YamlUnmarshal(consts.RedistoreConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
