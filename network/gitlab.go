package network

import "strings"

var (
	createdRunnerTokenPrefix = "glrt-"
)

func TokenIsCreatedRunnerToken(token string) bool {
	return strings.HasPrefix(token, createdRunnerTokenPrefix)
}
