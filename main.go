package main

import (
	"flag"
	"fmt"
)

func main() {

	msg := flag.String("message", "", "Message")
	_ = flag.String("importance", "low", "Importance: (low|normal|high)")
	flag.Parse()

	config, err := GetConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config)

	err = SendMessage(*msg, config.ChatId, config.Token)
	// fmt.Println(err)
}
