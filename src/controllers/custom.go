package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	config "webhook/src/config"
)

func CustomWebhook(w http.ResponseWriter, r *http.Request) {

	postBody, _ := json.Marshal(map[string]string{
		"content": "Teste",
	})

	requestBody := bytes.NewBuffer(postBody)
	resp, _ := http.Post(config.DiscordUrl, "application/json", requestBody)

	fmt.Fprintf(w, resp.Status)
}
