# sysinfobot

## A Telegram bot that provides information about the server where it resides.

This bot is still a work-in-progress, changes will be made, and currently doesn't reply with any useful information (unlike what the description says).

In the meantime, let me introduce you the configuration method.

### Installation

```shell
# get the source code
$ go get github.com/gsora/sysinfobot

# get the dependencies
$ go get github.com/fatih/structs

# build and install  
$ go build github.com/gsora/sysinfobot
$ go install github.com/gsora/sysinfobot

# ...write your configuration file and...
# ...run!
$ sysinfobot
```

These instructions assume you've correctly [installed and configured](https://golang.org/doc/install) a **Golang** environment.

### Configuration

You need to provide a valid JSON configuration file in `~/.config/sysinfobot.json`, or the bot will not operate.

You can find a sample configuration file in `support/configsample.conf` or at the end of this paragraph.

To secure the information provided, you need to create *your* personal Telegram bot, by following the [official documentation](https://core.telegram.org/bots#botfather).

You can create "bot commands" too, but it's not strictly needed.

The bot will obey to a well-defined set of commands that will be provided once it will be capable of at least work: remember, right now is **work-in-progress**!

```json
{
    "cert_path": "web server certificate",
    "key_path": "web server key",
    "url": "your server url, needs to be https://",
    "endpoint": "an endpoint to let telegram push new messages",
    "bot_token": "bot token",
    "port": "http port, telegram only allows 80, 88, 443, 8443",
    "authorized_users": [
        "list",
        "with",
        "autorized",
        "users"
    ]
}
```
