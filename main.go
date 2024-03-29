package main

import (
	"fmt"
	"log"
	"net/http"

	"webhook/src/config"
	controllers "webhook/src/controllers"

	"github.com/gorilla/mux"
)

func setupRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/webhook/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Ok") })
	r.HandleFunc("/webhook/custom", controllers.CustomWebhook)
	r.HandleFunc("/webhook/heroku", controllers.HerokuWebhook)
	r.HandleFunc("/webhook/vercel", controllers.VercelWebhook)
	r.HandleFunc("/webhook/financial", controllers.FinancialWebhook)

	return r
}

func main() {
	config.Load()

	log.Printf("Webhook running at the port :%d", config.Port)

	r := mux.NewRouter()
	r = setupRoutes(r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
