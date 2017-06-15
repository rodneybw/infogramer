# infogramer - Telegram Bot

## How to install

Requirements:
* Go >= 1.7.1
* Ubuntu/Debian (Only tested on these, but should work fine on any other distributions)

Follow these instructions to create the Telegram Bot and note down the Bot Token: https://core.telegram.org/bots#6-botfather Start a conversation with your bot and write him a message. **(Important for later steps!)**

Continue with these steps...
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

# Get the telegram chat id. Replace BOT_TOKEN with your token. You should see the message you sent earlier.
$ infogramer -getId -token="BOT_TOKEN"

# Example output: 
# Trying to get the telegram chat id...
# Message: 'one slice of pizza please'
# Is this message correct? Then your chat id is 123456789

# Create the config file. Replace CHAT_ID and BOT_TOKEN with yours.
$ echo '{"chat_id": CHAT_ID, "token": "BOT_TOKEN"}' > /etc/infogramer

# Test it :)
$ infogramer -message="Hello. Do you already feel informed?"
```

## Usage 
### Notification if a ssh login occurs on your machine

Add this line to `/etc/profile`:
```sh
infogramer -message="SSH Login @ $(hostname)\n$(date +%Y-%d-%m) - $(date +%H:%M)\nUser:$USER" # note: german time format ;)
```

### Good morning wishes using cron
Crontab:
```sh
...
0 9 * * * infogramer -message="Heeeeeeeey. A new day begins, sunshine. :)" >/dev/null 2>&1
...
```
