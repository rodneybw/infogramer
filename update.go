package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// This file is only used to get the chat id when installing the bot
// and shouldn't be used somewhere else.

type UpdateChat struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

type UpdateMessage struct {
	Chat UpdateChat `json:"chat"`
	Text string     `json:"text"`
}

type Update struct {
	Message UpdateMessage `json:"Message"`
}

type UpdateCollection struct {
	Updates []Update `json:"result"`
	Ok      bool
}

func GetChatId(token string) (int64, string, error) {

	url := "https://api.telegram.org/bot" + token + "/getUpdates"

	resp, err := http.Get(url)
	if err != nil {
		return 0, "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, "", err
	}

	coll := UpdateCollection{}

	err = json.Unmarshal(body, &coll)
	if err != nil {
		return 0, "", err
	}

	if len(coll.Updates) == 0 {
		return 0, "", errors.New("UpdateCollection is empty. Did you really send a message to the bot? (Also verify that your token is correct.)")
	}

	update := coll.Updates[len(coll.Updates)-1]

	return update.Message.Chat.Id, update.Message.Text, nil
}
