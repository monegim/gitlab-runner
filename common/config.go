package common

type Config struct {
	Runners    []*RunnerConfig
	Concurrent int
}

type RunnerCredentials struct {
	URL   string `toml:"url" env:"CI_SERVER_URL"`
	Token string
}

type RunnerConfig struct {
	RunnerCredentials
	RunnerSettings
}

type RunnerSettings struct {
}
