package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
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
	r.Get("/projects/{id}", env.getProject)
	r.Get("/projects/{id}/comments", env.getProjectComments)
	r.Post("/project", env.createProject)

	r.Post("/user", env.createUser)
	r.Get("/users", env.getUsers)

	r.Get("/payments", env.getPayments)
	r.Post("/payments", env.createPayment)

	r.Get("/categories", env.getCategories)
	r.Post("/category", env.createCategory)

	r.Get("/comments", env.getComments)
	r.Post("/comments", env.createComment)

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

func (env *Env) getProject(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "id")
	if projectID == "" {
		respondWithError(w, http.StatusNotFound, "No such project id")
		return
	}
	project, err := env.db.GetProject(projectID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, project)
}

func (env *Env) getProjectComments(w http.ResponseWriter, r *http.Request) {
	projectID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "No such project id")
		return
	}
	comments, err := env.db.AllProjectComments(projectID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, comments)
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

func (env *Env) getPayments(w http.ResponseWriter, r *http.Request) {
	users, err := env.db.AllPayments()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

type paymentRequest struct {
	UserID    int     `json:"userId"`
	ProjectID int     `json:"projectId"`
	Amount    float64 `json:"amount"`
}

func (env *Env) createPayment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payment paymentRequest
	err := decoder.Decode(&payment)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	paymentID, err := env.db.CreatePayment(
		payment.UserID,
		payment.ProjectID,
		payment.Amount,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, paymentID)
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

func (env *Env) getComments(w http.ResponseWriter, r *http.Request) {
	comments, err := env.db.AllComments()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, comments)
}

type commentRequest struct {
	UserID    int    `json:"userId"`
	ProjectID int    `json:"projectId"`
	Content   string `json:"content"`
}

func (env *Env) createComment(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var comment commentRequest
	err := decoder.Decode(&comment)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	commentID, err := env.db.CreateComment(
		comment.UserID,
		comment.ProjectID,
		comment.Content,
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, commentID)
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
