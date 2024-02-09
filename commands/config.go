package commands

import (
	"fmt"
	"simple-gitlab-runner/common"
	"sync"
)

type configOptions struct {
	config      *common.Config
	configMutex sync.Mutex
}

func (c *configOptions) RunnerByToken(token string) (*common.RunnerConfig, error) {
	config := c.getConfig()
	if config == nil {
		return nil, fmt.Errorf("config has not been loaded")
	}
	for _, runner := range config.Runners {
		if token == runner.Token {
			return runner, nil
		}
	}
	return nil, fmt.Errorf("could not find a runner with the token %s")
}

func (c *configOptions) getConfig() *common.Config {
	c.configMutex.Lock()
	defer c.configMutex.Unlock()

	if c.config == nil {
		return nil
	}

	config := *c.config
	return &config
}
