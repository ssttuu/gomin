package main

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/gomin/gomin"
	"github.com/stupschwartz/gomin/gomin/handler"
	"github.com/stupschwartz/gomin/gomin/minify"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	env := &handler.Env{
		DB: gomin.NewRedisClient(),
	}

	apiRouter := r.PathPrefix("/api/v0").Subrouter()

	minify.Register(apiRouter, env)

	log.Fatal(http.ListenAndServe(":8080", r))
}
