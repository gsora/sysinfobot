package main

import (
	"fmt"
	"github.com/gsora/sysinfobot/support"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var baseURL = ""

func main() {
	// check for presence and correctness of the configuration file
	conf, err := support.CheckConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	// set the base URL for api calls
	baseURL = "https://api.telegram.org/bot" + conf.BotToken + "/"

	// set the webhook as the configuration file says
	_, err = http.PostForm(baseURL+"setWebhook", url.Values{"url": {conf.URL + ":" + conf.Port + "/" + conf.Endpoint}})
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
	refChat := data.Message.Chat.ID
	sendMessage(strconv.Itoa(refChat), echoText)

}

// simple function to send a message back to its chat
func sendMessage(chatID string, text string) {
	params := url.Values{}
	params.Set("chat_id", chatID)
	params.Set("text", text)

	_, err := http.PostForm(baseURL+"sendMessage", params)
	if err != nil {
		log.Fatal(err)
	}
}
