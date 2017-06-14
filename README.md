# infogramer - Telegram Bot

## How to install

Requirements:
* Go >= 1.7.1

Follow these instructions to create the Telegram Bot: https://core.telegram.org/bots#6-botfather

```sh
# Download the repository
$ git clone https://github.com/rodneybw/infogramer.git

# Build the program
$ cd infogramer
$ go build

# Copy the executable to /usr/local/bin
$ cp ./infogramer /usr/local/bin/infogramer

# Make it executable
$ chmod a+x /usr/local/bin/infogramer

# Create the config file. Replace CHAT_ID with your chat id and BOT_TOKEN with the token botfather told you.
$ echo '{"chat_id": CHAT_ID, "token": "BOT_TOKEN"}' > /etc/infogramer

# Test it :)
$ infogramer -message="Hello. Do you already feel informed?"
```
    

