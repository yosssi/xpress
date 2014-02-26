package models

type Application struct {
	ServerConfig *ServerConfig
	LoggerConfig *LoggerConfig
}

func NewApplication() (*Application, error) {
	serverConfig, err := NewServerConfig()
	if err != nil {
		return nil, err
	}
	loggerConfig, err := NewLoggerConfig()
	if err != nil {
		return nil, err
	}
	app := &Application{ServerConfig: serverConfig, LoggerConfig: loggerConfig}
	return app, nil
}
