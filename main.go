package main

import (
	"fmt"
	"log"
	"net/http"

	"webhook/src/config"

	"github.com/gorilla/mux"
)

func setupRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Ok") })

	return r
}

func main() {
	config.Load()

	log.Printf("Webhook running at the port :%d", config.Port)

	r := mux.NewRouter()
	r = setupRoutes(r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
