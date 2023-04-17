package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webhook/src/config"
	"webhook/src/util"
)

type PayloadVercel struct {
	Job struct {
		Id        string `json:"id"`
		State     string `json:"state"`
		CreatedAt int64  `json:"createdAt"`
	} `json:"job"`
}

func VercelWebhook(w http.ResponseWriter, r *http.Request) {

	var payload PayloadVercel

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		util.CaptureException(err, nil)
		return
	}

	content := fmt.Sprintf(
		"```md\nVercel:\n=======\n[ID:](%s) [Status:](%s)\n[Criado em:](%s)```",
		payload.Job.Id, payload.Job.State, time.UnixMilli(payload.Job.CreatedAt).Local().Format(time.RFC822),
	)

	postBody, _ := json.Marshal(map[string]string{
		"content": content,
	})

	requestBody := bytes.NewBuffer(postBody)
	resp, _ := http.Post(config.DiscordUrl, "application/json", requestBody)

	fmt.Fprintf(w, resp.Status)
}
