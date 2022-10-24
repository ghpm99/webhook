package controllers

import (
	"bytes"
	"fmt"
	"encoding/json"
	"net/http"
	"webhook/src/config"
)

type PayloadCustom struct {
	Content string `json:"content"`
}

func CustomWebhook(w http.ResponseWriter, r *http.Request) {

	var payload PayloadCustom

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content := fmt.Sprintf(
		"```md\nCustom:\n=======\n[Content:](%s)```",
		payload.Content,
	)

	postBody, _ := json.Marshal(map[string]string{
		"content": content,
	})

	requestBody := bytes.NewBuffer(postBody)
	resp, _ := http.Post(config.DiscordUrl, "application/json", requestBody)

	fmt.Fprintf(w, resp.Status)
}
