package controllers

import (
	"fmt"
	"net/http"

	config "webhook/src/config"
)

type PayloadCustom struct {
	Content string `json:"content"`
}

func CustomWebhook(w http.ResponseWriter, r *http.Request) {

	resp, _ := http.Post(config.DiscordUrl, "application/json", r.Body)

	fmt.Fprintf(w, resp.Status)
}
