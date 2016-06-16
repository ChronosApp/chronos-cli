package main

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
)

type Remote struct {
	URL      string
	Usernme  string `json:"username"`
	password string `json:"password"`
}

func ParseConfig() (*Remote, error) {

	//obain user's homedir
	var Homedir string
	if u, err := user.Current(); err != nil {
		return nil, err
	} else {
		Homedir = u.HomeDir
	}

	config, err := ioutil.ReadFile(Homedir + "/.chronos")
	if err != nil {
		return nil, err
	}

	var server Remote

	err = json.Unmarshal(config, &server)
	if err != nil {
		return nil, err
	}

	return &server, nil
}
