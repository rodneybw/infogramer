package main

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"net/http"
)

type MessageParameters struct {
	ChatId int64  `json:"chat_id"` // telegram internal chat id
	Text   string `json:"text"`
}

type Message struct {
	MessageId int64 `json:"message_id"`
}

func SendMessage(text string, chatId int64, token string) error {

	url := "https://api.telegram.org/bot" + token + "/sendMessage"
	param := &MessageParameters{
		ChatId: chatId,
		Text:   text,
	}

	para, err := json.Marshal(param)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(para))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// fmt.Println(string(body))

	msg := Message{}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return err
	}

	return nil
}
