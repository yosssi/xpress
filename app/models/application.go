package models

import (
	"strconv"

	"github.com/yosssi/gold"
	"github.com/yosssi/gologger"
	"github.com/yosssi/xpress/app/consts"
)

// An Application represents an application context.
type Application struct {
	ServerConfig *ServerConfig
	LoggerConfig *LoggerConfig
	Logger       *gologger.Logger
	Generator    *gold.Generator
	Locale       string
	Dictionaries map[string]*Dictionary
}

// Port returns ServerConfig's Port.
func (a *Application) Port() int {
	return a.ServerConfig.Port
}

// PortString returns a string value of ServerConfig's Port.
func (a *Application) PortString() string {
	return strconv.Itoa(a.Port())
}

// Msg returns a message from dictionaries.
func (a *Application) Msg(s string) string {
	return a.Dictionaries[a.Locale].Msg(s)
}

// Development returns the Application's ServerConfig's Development.
func (a *Application) Development() bool {
	return a.ServerConfig.Development
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

	dictionaries, err := NewDictionaries([]string{consts.LocaleEn, consts.LocaleJa})
	if err != nil {
		return nil, err
	}

	return &Application{ServerConfig: serverConfig, LoggerConfig: loggerConfig, Logger: logger, Generator: generator, Locale: consts.LocaleEn, Dictionaries: dictionaries}, nil
}
