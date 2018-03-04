package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/message", handleQryMessage).Methods("GET")
	router.HandleFunc("/m/{msg}", handleUrlMessage).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func handleQryMessage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func handleUrlMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}
