package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App for initial db and router setup
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user string, password string, dbname string) {}

func (a *App) Run(addr string) {}
