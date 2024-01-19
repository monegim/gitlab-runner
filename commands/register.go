package commands

import (
	"bufio"
	"os"
)

type RegisterCommand struct {
	url    string
	token  string
	reader *bufio.Reader
}

func Register() {
	rc := RegisterCommand{}
	rc.askRunner()
}

func (rc *RegisterCommand) askRunner() {
	rc.url = rc.ask("url", "Please Enter url:")
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
	if output != "" {
		*result = output
		return true
	}
	if allowEmpty {
		return true
	}
	return false
}
