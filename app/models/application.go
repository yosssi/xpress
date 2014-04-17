package models

import (
	"os"
	"strconv"

	"github.com/boj/redistore"
	"github.com/gorilla/securecookie"
	"github.com/yosssi/goelasticsearch"
	"github.com/yosssi/gogithub"
	"github.com/yosssi/gold"
	"github.com/yosssi/gologger"
	"github.com/yosssi/xpress/app/consts"
)

// An Application represents an application context.
type Application struct {
	ServerConfig        *ServerConfig
	LoggerConfig        *LoggerConfig
	ElasticsearchConfig *ElasticsearchConfig
	RediStoreConfig     *RediStoreConfig
	Logger              *gologger.Logger
	Generator           *gold.Generator
	Locale              string
	Dictionary          *Dictionary
	GitHubClient        *gogithub.Client
	ElasticsearchClient *goelasticsearch.Client
	RediStoreKeyPair    []byte
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
	return a.Dictionary.Msg(s)
}

// Development returns the Application's ServerConfig's Development.
func (a *Application) Development() bool {
	return a.ServerConfig.Development
}

// NewRediStore generates RediStore and returns it.
func (a *Application) NewRediStore() (*redistore.RediStore, error) {
	return redistore.NewRediStore(a.RediStoreConfig.Size, a.RediStoreConfig.Network, a.RediStoreConfig.Address, a.RediStoreConfig.Password, a.RediStoreKeyPair)
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

	elasticsearchConfig, err := NewElasticsearchConfig()
	if err != nil {
		return nil, err
	}

	rediStoreConfig, err := NewRediStoreConfig()
	if err != nil {
		return nil, err
	}

	logger := &gologger.Logger{Name: loggerConfig.Name, Level: loggerConfig.Level, File: loggerConfig.File}

	generator := gold.NewGenerator(!serverConfig.Development)

	locale := consts.LocaleEn

	dictionary, err := NewDictionary(locale)
	if err != nil {
		return nil, err
	}

	githubClient := gogithub.NewClient(os.Getenv(consts.EnvGitHubClientID), os.Getenv(consts.EnvGitHubClientSecret))

	elasticsearchClient := goelasticsearch.NewClient(elasticsearchConfig.BaseURL)

	return &Application{ServerConfig: serverConfig, LoggerConfig: loggerConfig, ElasticsearchConfig: elasticsearchConfig, RediStoreConfig: rediStoreConfig, Logger: logger, Generator: generator, Locale: locale, Dictionary: dictionary, GitHubClient: githubClient, ElasticsearchClient: elasticsearchClient, RediStoreKeyPair: []byte(securecookie.GenerateRandomKey(32))}, nil
}
