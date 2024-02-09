package common

import "github.com/sirupsen/logrus"

var executorProviders map[string]ExecutorProvider
func GetExecuteNames() []string {
	var names []string
	for name := range executorProviders {
		names = append(names, name)
	}

	return names
}

type ExecutorProvider interface {

}

func RegisterExecutorProvider(executor string, provider ExecutorProvider) {
	logrus.Debugln("Registering", executor, "executor...")
}