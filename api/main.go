package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"api/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Env for persisting states global to server
type Env struct {
	db models.Datastore
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.Recoverer)

	db, err := models.NewDB()
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db}
	r.Get("/projects", env.getProjects)

	port := os.Getenv("PORT")
	log.Println("Running server at port " + port)
	http.ListenAndServe(":"+port, r)
}

func (env *Env) getProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := env.db.AllProjects()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, projects)
}

/*
// Initialize sets up the database connection and routes for the server
func (server *Server) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=postgres sslmode=disable user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	// Try to connect every second
	for i := 0; i < 15; i++ {
		server.DB, err = sql.Open("postgres", connectionString)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
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
	server.Router.HandleFunc("/projects", server.getProjects).Methods("GET")
	server.Router.HandleFunc("/project", server.createProject).Methods("POST")
	server.Router.HandleFunc("/project/{id:[0-9]+}", server.getProject).Methods("GET")
	server.Router.HandleFunc("/project/{id:[0-9]+}", server.updateProject).Methods("PUT")
	server.Router.HandleFunc("/project/{id:[0-9]+}", server.deleteProject).Methods("DELETE")
}

func (server *Server) getProjects(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	projects, err := getProjects(server.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, projects)
}

func (server *Server) createProject(w http.ResponseWriter, r *http.Request) {
	var p project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.createProject(server.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (server *Server) getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid project ID")
		return
	}

	p := project{ID: id}
	if err := p.getProject(server.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Project not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (server *Server) updateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid project ID")
		return
	}

	var p project
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := p.updateProject(server.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (server *Server) deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Project ID")
		return
	}

	p := project{ID: id}
	if err := p.deleteProject(server.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}


*/

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
