package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Server for initial db and router setup
type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {
	server := Server{}
	server.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(port)
}

// Initialize sets up the database connection and routes for the server
func (server *Server) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=postgres sslmode=disable user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	server.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = server.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// Run starts the server and serves on the specified port
func (server *Server) Run(port string) {
	server.Router = mux.NewRouter()
	server.initializeRoutes()

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	corsHandler := handlers.CORS(originsOk, headersOk, methodsOk)(server.Router)
	fmt.Println("Running server at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	server.Router.HandleFunc("/products", server.getProducts).Methods("GET")
	server.Router.HandleFunc("/product", server.createProduct).Methods("POST")
	server.Router.HandleFunc("/product/{id:[0-9]+}", server.getProduct).Methods("GET")
	server.Router.HandleFunc("/product/{id:[0-9]+}", server.updateProduct).Methods("PUT")
	server.Router.HandleFunc("/product/{id:[0-9]+}", server.deleteProduct).Methods("DELETE")
}

func (server *Server) getProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getProducts(server.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (server *Server) createProduct(w http.ResponseWriter, r *http.Request) {
	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.createProduct(server.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (server *Server) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := product{ID: id}
	if err := p.getProduct(server.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (server *Server) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := p.updateProduct(server.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (server *Server) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := product{ID: id}
	if err := p.deleteProduct(server.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!")
}
