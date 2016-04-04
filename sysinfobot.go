package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gsora/sysinfobot/support"
)

var baseURL = ""
var conf support.ConfigFile
var err error

func main() {
	// check for presence and correctness of the configuration file
	conf, err = support.CheckConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	// set the base URL for api calls
	baseURL = "https://api.telegram.org/bot" + conf.BotToken + "/"

	// set the webhook as the configuration file says
	_, err = http.PostForm(baseURL+"/setWebhook", url.Values{"url": {conf.URL + ":" + conf.Port + "/" + conf.Endpoint}})
	if err != nil {
		log.Fatal(err)
	}

	// start the bot!
	fmt.Println("--> Starting sysinfobot")
	err = support.PrintBotInformations(conf.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/"+conf.Endpoint, endpointHandler)
	log.Fatal(http.ListenAndServeTLS(":"+conf.Port, conf.CertPath, conf.KeyPath, nil))
}

// handle the webhook data
func endpointHandler(w http.ResponseWriter, r *http.Request) {
	data := support.LoadJSONToTelegramObject(r.Body)
	echoText := data.Message.Text
	secureSendMessage(data, echoText)

}

// simple function to send a message back to its chat, and check for security
func secureSendMessage(tObj support.TelegramObject, text string) {

	recipient := tObj.Message.From.Username

	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(tObj.Message.Chat.ID))
	params.Set("text", "Not authorized.")

	for _, username := range conf.AuthorizedUsers {
		if username == recipient {
			params.Del("text")
			params.Set("text", text)
			break
		}
	}

	_, err := http.PostForm(baseURL+"sendMessage", params)
	if err != nil {
		log.Fatal(err)
	}
}
