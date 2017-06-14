package main

import (
	"flag"
	"fmt"
)

func main() {

	// call like that: 'infogramer --message="Hello. How are you? Do you feel informed?"'
	msg := flag.String("message", "", "Message")
	flag.Parse()

	// Configuration contains the bot token and the chat id (see: https://core.telegram.org/api)
	config, err := GetConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = SendMessage(*msg, config.ChatId, config.Token)
}
