package models

import (
	"github.com/yosssi/gold"
	"github.com/yosssi/gologger"
	"strconv"
)

// An Application represents an application context.
type Application struct {
	ServerConfig *ServerConfig
	LoggerConfig *LoggerConfig
	Logger       *gologger.Logger
	Generator    *gold.Generator
}

// Port returns ServerConfig's Port.
func (a *Application) Port() int {
	return a.ServerConfig.Port
}

// PortString returns a string value of ServerConfig's Port.
func (a *Application) PortString() string {
	return strconv.Itoa(a.Port())
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

	generator := gold.NewGenerator(!serverConfig.Development)

	return &Application{ServerConfig: serverConfig, LoggerConfig: loggerConfig, Logger: logger, Generator: generator}, nil
}
