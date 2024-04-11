package oauth

const (
	// RequestFailed is the error message when a request fails.
	RequestFailed = "request failed"
	// ResponseFailed is the error message when a response fails.
	ResponseFailed = "response failed"
)

type GlobalConfig struct {
	CallbackURL string
	FrontendURL string
}

var (
	githubConfig  GithubConfig  //nolint:gochecknoglobals // This is a configuration object
	discordConfig DiscordConfig //nolint:gochecknoglobals // This is a configuration object
	oauthConfig   GlobalConfig  //nolint:gochecknoglobals // This is a configuration object
)

type GithubConfig struct {
	ClientID     string
	ClientSecret string
}

type DiscordConfig struct {
	ClientID     string
	ClientSecret string
	URL          string
}

func GetGithubClientID() string {
	return githubConfig.ClientID
}

func GetCallbackURL() string {
	return oauthConfig.CallbackURL
}

func GetFrontendURL() string {
	return oauthConfig.FrontendURL
}

func GetDiscordURL() string {
	return discordConfig.URL
}

// InitGithub initializes the Github oauth configuration.
func InitGithub(config GithubConfig) {
	githubConfig = config
}

// InitDiscord initializes the Discord oauth configuration.
func InitDiscord(config DiscordConfig) {
	discordConfig = config
}

func SetOauthConfig(config GlobalConfig) {
	oauthConfig = config
}
