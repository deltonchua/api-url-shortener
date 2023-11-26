package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/deltonchua/api-url-shortener/postgres/store"
	"github.com/go-chi/chi/v5"
)

func (s *server) handleStatus() http.HandlerFunc {
	type res struct {
		Status    string    `json:"status"`
		Timestamp time.Time `json:"timestamp"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.sendJSON(w, res{
			Status:    "ok",
			Timestamp: time.Now().UTC().Truncate(time.Microsecond),
		})
	}
}

func (s *server) handleShortenURL() http.HandlerFunc {
	type req struct {
		Url string `json:"url"`
	}
	type res struct {
		ID string `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var body req
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			s.sendError(w, http.StatusBadRequest)
			return
		}
		if !isURL(body.Url) {
			s.sendError(w, http.StatusBadRequest)
			return
		}
		tx, err := s.db.Begin()
		if err != nil {
			s.sendInternalError(w, err)
			return
		}
		defer tx.Rollback()
		qtx := s.queries.WithTx(tx)
		ctx := context.Background()
		publicID, err := qtx.GetID(ctx, body.Url)
		if err != nil && err != sql.ErrNoRows {
			s.sendInternalError(w, err)
			return
		}
		if err == sql.ErrNoRows {
			publicID, err = nanoid()
			if err != nil {
				s.sendInternalError(w, err)
				return
			}
			if err := qtx.CreateURL(ctx, store.CreateURLParams{PublicID: publicID, Url: body.Url}); err != nil {
				s.sendInternalError(w, err)
				return
			}
			if err := tx.Commit(); err != nil {
				s.sendInternalError(w, err)
				return
			}
		}
		s.sendJSON(w, res{publicID})
	}
}

func (s *server) handleGetURL() http.HandlerFunc {
	type res struct {
		Url string `json:"url"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if len(id) != 8 {
			s.sendError(w, http.StatusNotFound)
			return
		}
		tx, err := s.db.Begin()
		if err != nil {
			s.sendInternalError(w, err)
			return
		}
		defer tx.Rollback()
		qtx := s.queries.WithTx(tx)
		ctx := context.Background()
		data, err := qtx.GetURL(ctx, id)
		if err != nil && err != sql.ErrNoRows {
			s.sendInternalError(w, err)
			return
		}
		if err == sql.ErrNoRows {
			s.sendError(w, http.StatusNotFound)
			return
		}
		if err := qtx.UpdateCount(ctx, store.UpdateCountParams{PublicID: id, Count: data.Count + 1}); err != nil {
			s.sendInternalError(w, err)
			return
		}
		if err := tx.Commit(); err != nil {
			s.sendInternalError(w, err)
			return
		}
		s.sendJSON(w, res{data.Url})
	}
}

func (s *server) handleGetCount() http.HandlerFunc {
	type res struct {
		Count int64 `json:"count"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if len(id) != 8 {
			s.sendError(w, http.StatusNotFound)
			return
		}
		count, err := s.queries.GetCount(context.Background(), id)
		if err != nil && err != sql.ErrNoRows {
			s.sendInternalError(w, err)
			return
		}
		if err == sql.ErrNoRows {
			s.sendError(w, http.StatusNotFound)
		}
		s.sendJSON(w, res{count})
	}
}
