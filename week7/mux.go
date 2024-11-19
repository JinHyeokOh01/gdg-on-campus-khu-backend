package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/clock"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/config"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/handler"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error){
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json; charset=stf-8")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg) // connection 생성
	if err != nil{
		return nil, cleanup, err
	}
	r := store.Repository(Clocker: clock.RealClocker{})
	at := &handler.AddTask{DB: db, Repo: &r, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{DB: db, Repo: &r}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux, cleanup, nil
}