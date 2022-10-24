package controllers

import (
	"bytes"
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"webhook/src/config"
)

type myTime time.Time

var _ json.Unmarshaler = &myTime{}

func (mt *myTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = myTime(t)
	return nil
}

type PayloadFinancial struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	PaymentDate *myTime `json:"payment_date"`
	Value float32 `json:"value"`
	Payment string `json:"payment"`
}

func FinancialWebhook(w http.ResponseWriter, r *http.Request) {

	var payload PayloadFinancial

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	content := fmt.Sprintf(
		"```md\nFinanceiro:\n=======\n[Id:](%s) [Tipo:](%s)\n[Name:](%s)\n[Data de Pagamento:](%t) [Valor:](%f)\n[Pagamento Link:](%s)```",
		payload.Id, payload.Type, payload.Name, payload.PaymentDate, payload.Value, payload.Payment,
	)

	postBody, _ := json.Marshal(map[string]string{
		"content": content,
	})

	requestBody := bytes.NewBuffer(postBody)
	resp, _ := http.Post(config.DiscordUrl, "application/json", requestBody)

	fmt.Fprintf(w, resp.Status)
}
