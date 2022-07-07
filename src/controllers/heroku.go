package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webhook/src/config"
)

type Payload struct {
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
		GitUrl      string    `json:"git_url"`
		Maintenance bool      `json:"maintenance"`
		Name        string    `json:"name"`
		WebUrl      string    `json:"web_url"`
	} `json:"data"`
}

func HerokuWebhook(w http.ResponseWriter, r *http.Request) {

	var action Payload

	err := json.NewDecoder(r.Body).Decode(&action)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content := fmt.Sprintf("Heroku:\nAção: %s Autor: %s Criado em: %s\nId: %s Url repositorio: %s App Nome: %s\nEm manutenção: %t App Url: %s", action.Action, action.Actor.Email, action.Data.CreatedAt, action.Data.Id, action.Data.GitUrl, action.Data.Name, action.Data.Maintenance, action.Data.WebUrl)

	postBody, _ := json.Marshal(map[string]string{
		"content": content,
	})

	requestBody := bytes.NewBuffer(postBody)
	resp, _ := http.Post(config.DiscordUrl, "application/json", requestBody)

	fmt.Fprintf(w, resp.Status)

}
