package models

import (
	"github.com/yosssi/gologger"
	"strconv"
)

// An Application represents an application context.
type Application struct {
	ServerConfig *ServerConfig
	LoggerConfig *LoggerConfig
	Logger       *gologger.Logger
}

// PortString returns a string value of ServerConfig's Port.
func (a *Application) PortString() string {
	return strconv.Itoa(a.ServerConfig.Port)
}

// NewApplication generates an Application and returns it.
func NewApplication() (*Application, error) {
	serverConfig, err := NewServerConfig()
	if err != nil {
		return nil, err
	}

	loggerConfig, err := NewLoggerConfig()
	if err != nil {
		return nil, err
	}

	logger := &gologger.Logger{Name: loggerConfig.Name, Level: loggerConfig.Level, File: loggerConfig.File}

	app := &Application{ServerConfig: serverConfig, LoggerConfig: loggerConfig, Logger: logger}

	return app, nil
}
