package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/deltonchua/api-url-shortener/postgres/store"
	"github.com/go-chi/chi/v5"
	// "github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type server struct {
	mux     *chi.Mux
	db      *sql.DB
	queries *store.Queries
}

func Run() error {
	// if err := godotenv.Load(); err != nil {
	// 	return err
	// }
	s := newServer()
	db, queries, err := connectDB()
	if err != nil {
		return err
	}
	s.db = db
	s.queries = queries
	return http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), s.mux)
}

func newServer() *server {
	s := &server{
		mux: chi.NewRouter(),
	}
	s.routes()
	return s
}

func connectDB() (*sql.DB, *store.Queries, error) {
	connStr := fmt.Sprintf("host=%v user=%v password=%v dbname=%v sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, nil, err
	}
	return db, store.New(db), nil
}
