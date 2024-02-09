/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// import "simple-gitlab-runner/cmd"

// func main() {
// 	cmd.Execute()
// }

type Config struct {
	Concurrent int64  `toml:"concurrent"`
	URL        string `toml:"url"`
	Token      string `toml:"token"`
}

func main() {
	var config Config
	path := "./.config.toml"
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		logrus.Fatalln("decoding file", err)
	}
	isRoot()
	new_file := "new_config.toml"
	if _, err := os.Stat(new_file); err != nil {
		os.Create(new_file)
	}
	of, err := os.OpenFile(new_file, os.O_CREATE|os.O_RDWR, 0755)
	defer of.Close()
	if err != nil {
		logrus.Fatalln("opening file:", err)
	}
	en := toml.NewEncoder(of)
	err = en.Encode(config)
	if err != nil {
		logrus.Errorln("encode problem:", err)
	}
}

func isRoot() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, err
	}
	return currentUser.Username == "root", nil
}
