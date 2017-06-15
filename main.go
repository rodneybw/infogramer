package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {

	// call like that: 'infogramer --message="Hello. How are you? Do you feel informed?"'
	msg := flag.String("message", "", "Message")
	getId := flag.Bool("getId", false, "Get the chat id. (Use it only during the installation steps. Message will be ignored.)")
	token := flag.String("token", "", "Telegram Bot Token. Only use this parameter if you use getId during the installation process.")
	flag.Parse()

	// get the telegram chat id (during the installation process)
	if *getId {

		if len(*token) == 0 {
			err := errors.New("You have to provide the telegram bot token.")
			fmt.Println(err)
			return
		}

		chatId, text, err := GetChatId(*token)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Trying to get the telegram chat id...")
		fmt.Printf("\nMessage: '%s'\n\n", text)
		fmt.Printf("Is this message correct? Then your chat id is %d\n\n", chatId)

		return // message will be ignored
	}

	// Configuration contains the bot token and the chat id (see: https://core.telegram.org/api)
	config, err := GetConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = SendMessage(*msg, config.ChatId, config.Token)
}
