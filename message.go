package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type MessageParameters struct {
	ChatId    int64  `json:"chat_id"` // telegram internal chat id
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"` // html or markdown
	Silent    bool   `json:"disable_notification"`
}

type Message struct {
	MessageId int64 `json:"message_id"`
}

func SendMessage(text string, chatId int64, silent bool, token string) error {

	// interpret special chars similar to echo (see: http://linuxcommand.org/lc3_man_pages/echoh.html)
	repl := strings.NewReplacer("\\n", "\n", "\\t", "\t")
	text = repl.Replace(text)

	url := "https://api.telegram.org/bot" + token + "/sendMessage"
	param := &MessageParameters{
		ChatId:    chatId,
		Text:      text,
		ParseMode: "markdown", // see: https://core.telegram.org/bots/api#markdown-style
		Silent:    silent,
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
