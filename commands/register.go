package commands

import (
	"bufio"
	"os"
	"simple-gitlab-runner/common"
	"simple-gitlab-runner/network"
	"strings"

	"github.com/sirupsen/logrus"
)

type RegisterCommand struct {
	reader *bufio.Reader
	configOptions
	Description string //	no	Description of the runner
	// info	hash //	no	Runner’s metadata. You can include name, version, revision, platform, and architecture, but only version, platform, and architecture are displayed in the Admin Area of the UI
	// active	bool //	no	Deprecated: Use paused instead. Specifies if the runner is allowed to receive new jobs
	Paused          bool   //	no	Specifies if the runner should ignore new jobs
	Locked          bool   //	no	Specifies if the runner should be locked for the current project
	RunUntagged     bool   //	no	Specifies if the runner should handle untagged jobs
	TagList         string // array	no	A list of runner tags
	AccessLevel     string //	no	The access level of the runner; not_protected or ref_protected
	MaximumTimeout  int    //	no	Maximum timeout that limits the amount of time (in seconds) that runners can run jobs
	MaintainerNote  string //	no	Deprecated, see maintenance_note
	MaintenanceNote string //	no	Free-form maintenance notes for the runner (1024 characters)
	common.RunnerCredentials
}

func Register() {
	rc := RegisterCommand{}
	rc.askRunner()
}

func (rc *RegisterCommand) askRunner() {
	rc.URL = rc.ask("url", "Please Enter url:")
	if rc.Token == "" {
		rc.Token = rc.ask("registration-token", "Enter the registration token:")
	}
	if !rc.tokenIsRunnerToken() {
		//	Do legacy registration
	}
	if r, err := rc.RunnerByToken(rc.Token); err == nil && r != nil {
		logrus.Warningln("A runner with this system ID and token has already been registered.")
	}
	rc.verifyRunner()

}

func (rc *RegisterCommand) ask(key, prompt string, allowEmpty ...bool) string {
	allowEmptyBool := len(allowEmpty) > 0 && allowEmpty[0]
	var result string
	for {
		if rc.askOnce(prompt, &result, allowEmptyBool) {
			break
		}
	}
	return result
}

func (rc *RegisterCommand) askOnce(prompt string, result *string, allowEmpty bool) bool {
	if rc.reader == nil {

		rc.reader = bufio.NewReader(os.Stdin)
	}
	line, _, err := rc.reader.ReadLine()
	if err != nil {
		panic(err)
	}
	output := string(line)
	output = strings.TrimSpace(output)
	if output != "" {
		*result = output
		return true
	}
	if allowEmpty {
		return true
	}
	return false
}

func (rc *RegisterCommand) tokenIsRunnerToken() bool {
	return network.TokenIsCreatedRunnerToken(rc.token)
}
