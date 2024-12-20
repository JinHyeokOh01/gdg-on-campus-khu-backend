package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/clock"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/config"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/handler"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/service"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/store"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week8/auth"
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
	clocker := clock.RealClocker{}
	r := store.Repository{Clocker: clocker}
	rcli, err := store.NewKVS(ctx, cfg)
	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil{
		return nil, cleanup, err
	}
	l := &handler.Login{
		Service: &service.Login{
			DB:				db,
			Repo:			&r,
			TokenGenerator:	jwter,
		},
		Validator: v,
	}
	mux.Post("/login", l.ServeHTTP)
	at := &handler.AddTask{
		Service: &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	//mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db, Repo: &r},
	}
	//mux.Get("/tasks", lt.ServeHTTP)
	mux.Route("/tasks", func(r chi.Router){
		r.Use(handler.AuthMiddleware(jwter))
		r.Post("/", at.ServeHTTP)
		r.Get("/", lt.ServeHTTP)
	})
	//추가
	mux.Route("/admin", func(r chi.Router){
		r.Use(handler.AuthMiddleware(jwter), handler.AdminMiddleware)
		r.Get("/", func(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, _ = w.Write([]byte(`{"message": "admin only"}`))
		})
	})
	ru := &handler.RegisterUser{
		Service: &service.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register", ru.ServeHTTP)
	return mux, cleanup, nil
}