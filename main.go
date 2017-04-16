package main

import (
	"flag"
	"fmt"
)

func main() {

	message := flag.String("message", "", "Message")
	importance := flag.String("importance", "low", "Importance: (low|normal|high)")
	flag.Parse()

	fmt.Printf("Message: %s\nImportance: %s\n", *message, *importance)
}
