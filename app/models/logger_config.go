package models

import (
	"github.com/yosssi/xpress/app/consts"
	"github.com/yosssi/xpress/app/utils"
)

// A LoggerConfig represents a logger config.
type LoggerConfig struct {
	Name  string `yaml:"name"`
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

// NewLoggerConfig parses a yaml file, generates a LoggerConfig and returns it.
func NewLoggerConfig() (*LoggerConfig, error) {
	config := &LoggerConfig{}
	if err := utils.YamlUnmarshal(consts.LoggerConfigPath, config); err != nil {
		return nil, err
	}
	return config, nil
}
