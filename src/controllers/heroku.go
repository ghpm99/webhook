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

type PayloadHeroku struct {
	Action string `json:"action"`
	Actor  struct {
		Email string `json:"email"`
		Id    string `json:"id"`
	} `json:"actor"`
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"id"`
	Data      struct {
		CreatedAt   time.Time `json:"created_at"`
		Id          string    `json:"id"`
		Maintenance bool      `json:"maintenance"`
		Name        string    `json:"name"`
		WebUrl      string    `json:"web_url"`
		App         struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"app"`
		Release struct {
			Id      string `json:"id"`
			Version int    `json:"version"`
		} `json:"release"`
	} `json:"data"`
}

func HerokuWebhook(w http.ResponseWriter, r *http.Request) {

	var payload PayloadHeroku

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		util.CaptureException(err, nil)
		return
	}

	content := fmt.Sprintf(
		"```md\nHeroku:\n=======\n[Ação:](%s) [Autor:](%s)\n[Criado em:](%s)\n[Id:](%s)\n[App Nome:](%s)\n[Em manutenção:](%t)\n[Release:](%d)```",
		payload.Action, payload.Actor.Email, payload.Data.CreatedAt, payload.Data.Id, payload.Data.App.Name, payload.Data.Maintenance,
		payload.Data.Release.Version,
	)

	postBody, _ := json.Marshal(map[string]string{
		"content": content,
	})

	requestBody := bytes.NewBuffer(postBody)
	resp, _ := http.Post(config.DiscordUrl, "application/json", requestBody)

	fmt.Fprintf(w, resp.Status)

}
