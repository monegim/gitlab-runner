package common

import "time"

type VerifyRunnerResponse struct {
	ID             int64     `json:"id,omitempty"`
	Token          string    `json:"token,omitempty"`
	TokenExpiresAt time.Time `json:"token_expire_at,omitempty"`
}
type Network interface {
	VerifyRunner(config RunnerCredentials, systemID string) *VerifyRunnerResponse
}
