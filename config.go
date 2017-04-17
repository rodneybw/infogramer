package main

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
)

const filename string = "/etc/infogramer"

type Config struct {
	Token  string `json:"token"`
	ChatId int64  `json:"chat_id"`
}

func GetConfig() (*Config, error) {

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
