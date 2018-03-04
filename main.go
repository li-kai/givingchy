package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App for initial db and router setup
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/message", handleQryMessage).Methods("GET")
	router.HandleFunc("/m/{msg}", handleURLMessage).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	corsHandler := handlers.CORS(originsOk, headersOk, methodsOk)(router)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
	fmt.Println("Running server at port" + port)

	fmt.Println(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))
}

func handleQryMessage(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	message := vars.Get("msg")

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func handleURLMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := vars["msg"]

	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}
