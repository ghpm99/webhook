package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webhook/src/config"
	"webhook/src/util"
)

type PayloadFinancial struct {
	Data []PaymentPayload `json:"data"`
}

type PaymentPayload struct {
	Id          int     `json:"id"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	PaymentDate string  `json:"payment_date"`
	Value       float32 `json:"value"`
	Payment     string  `json:"payment"`
}

func FinancialWebhook(w http.ResponseWriter, r *http.Request) {

	var payload PayloadFinancial

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		util.CaptureException(err, nil)
		return
	}

	content := ""

	for _, payment := range payload.Data {
		paymentEmbed := fmt.Sprintf(
			"```md\nFinanceiro:\n=======\n[Id:](%v) [Tipo:](%s)\n[Name:](%s)\n[Data de Pagamento:](%s) [Valor:](%s)\n```[Pagamento Link](%s)",
			payment.Id, payment.Type, payment.Name, payment.PaymentDate, fmt.Sprintf("R$%.2f", payment.Value), payment.Payment,
		)

		lengthContent := len(content)
		lengthPayment := len(paymentEmbed)

		if lengthContent >= 2000 || lengthContent+lengthPayment >= 2000 {
			postBody, _ := json.Marshal(map[string]string{
				"content": content,
			})
			requestBody := bytes.NewBuffer(postBody)
			http.Post(config.DiscordUrl, "application/json", requestBody)
			content = ""
		}
		content += paymentEmbed
	}

	postBody, _ := json.Marshal(map[string]string{
		"content": content,
	})

	requestBody := bytes.NewBuffer(postBody)
	resp, _ := http.Post(config.DiscordUrl, "application/json", requestBody)

	fmt.Fprintf(w, resp.Status)
}
