package common

import "time"

type Config struct {
	Runners    []*RunnerConfig
	Concurrent int
}

type RunnerCredentials struct {
	URL             string    `toml:"url" env:"CI_SERVER_URL"`
	Token           string    `toml:"token" json:"token"`
	ID              int64     `toml:"id" json:"id"`
	TokenObtainedAt time.Time `toml:"token_obtained_at" json:"token_obtained_at" description:"When the runner authentication token was obtained"`
	TokenExpiresAt  time.Time `toml:"token_expires_at" json:"token_expires_at" description:"Runner token expiration time"`
}

type RunnerConfig struct {
	Name string `toml:"name" json:"name"`
	RunnerCredentials
	RunnerSettings

	SystemIDState *SystemIDState `toml:"-" json:",omitempty"`
}

type RunnerSettings struct {
}
