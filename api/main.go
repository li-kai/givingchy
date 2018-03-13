package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

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
	r.Post("/project", env.createProject)
	r.Post("/user", env.createUser)
	r.Get("/users", env.getUsers)
	r.Get("/categories", env.getCategories)
	r.Post("/category", env.createCategory)

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

type projectRequest struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	AmountRequired float64   `json:"amountRequired"`
	EndTime        time.Time `json:"endTime"`
	Category       string    `json:"category"`
	UserID         int       `json:"userId"`
}

func (env *Env) createProject(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var project projectRequest
	err := decoder.Decode(&project)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	projectID, err := env.db.CreateProject(
		project.Title,
		project.Description,
		project.AmountRequired,
		project.EndTime,
		project.Category,
		project.UserID,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, projectID)
}

func (env *Env) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := env.db.AllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

type userRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (env *Env) createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u userRequest
	err := decoder.Decode(&u)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	userID, err := env.db.CreateUser(u.Email, u.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, userID)
}

func (env *Env) getCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := env.db.AllCategories()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, categories)
}

func (env *Env) createCategory(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var category models.Category
	err := decoder.Decode(&category)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = env.db.CreateCategory(category.Name)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, category)
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
