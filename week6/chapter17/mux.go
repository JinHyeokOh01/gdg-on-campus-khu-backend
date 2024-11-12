package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week6/chapter17/handler"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week6/chapter17/store"
)

func NewMux() http.Handler{
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json; charset=stf-8")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
	v := validator.New()
	mux.Handle("/tasks", &handler.AddTask{Store: store.Tasks, Validator: v})
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux
}