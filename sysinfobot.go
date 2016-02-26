package main

import (
	"fmt"
	"github.com/gsora/sysinfobot/support"
	"log"
	"net/http"
	"net/url"
)

func main() {
	conf, err := support.CheckConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	baseURL := "https://api.telegram.org/bot" + conf.BotToken + "/"

	//declare eventHook for method messageUpdate
	fmt.Println(conf.URL + ":" + conf.Port + "/" + conf.Endpoint)
	_, err = http.PostForm(baseURL+"setWebhook", url.Values{"url": {conf.URL + ":" + conf.Port + "/" + conf.Endpoint}})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/"+conf.Endpoint, endpointHandler)
	log.Fatal(http.ListenAndServeTLS(":"+conf.Port, conf.CertPath, conf.KeyPath, nil))
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	data := support.LoadJSONToTelegramObject(r.Body)
	fmt.Println(data)
}
