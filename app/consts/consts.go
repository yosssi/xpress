package consts

const (
	YmlExtension                   = ".yml"
	ConfigPath                     = "./config/"
	DictionariesPath               = "./dictionaries/"
	ServerConfigPath               = ConfigPath + "server" + YmlExtension
	LoggerConfigPath               = ConfigPath + "logger" + YmlExtension
	ElasticsearchConfigPath        = ConfigPath + "elasticsearch" + YmlExtension
	RediStoreConfigPath            = ConfigPath + "redistore" + YmlExtension
	LocaleEn                       = "en"
	LocaleJa                       = "ja"
	EnvGitHubClientID              = "X_GH_ID"
	EnvGitHubClientSecret          = "X_GH_SECRET"
	ElasticsearchIndexXpress       = "xpress"
	ElasticsearchTypeUser          = "user"
	SessionKeyUserID               = "user_id"
	ERR_MSG_SECURECOOKIE_NOT_VALID = "securecookie: the value is not valid"
)
