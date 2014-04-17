package models

import (
	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/utils"
)

// A RediStoreConfig represents a Redistore config.
type RediStoreConfig struct {
	Size       int    `yaml:"size"`
	Network    string `yaml:"network"`
	Address    string `yaml:"address"`
	Password   string `yaml:"password"`
	MaxAge     int    `yaml:"max_age"`
	SessionKey string `yaml:"session_key"`
}

// NewRedistoreConfig parses a yaml file, generates a RediStoreConfig and returns it.
func NewRediStoreConfig() (*RediStoreConfig, error) {
	config := &RediStoreConfig{}
	if err := utils.YamlUnmarshal(consts.RediStoreConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
