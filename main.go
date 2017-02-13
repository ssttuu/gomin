package main

import (
	"github.com/gorilla/mux"
	"github.com/stupschwartz/gomin/gomin"
	"github.com/stupschwartz/gomin/gomin/env"
	"github.com/stupschwartz/gomin/gomin/minify"
	"log"
	"net/http"
	"github.com/stupschwartz/gomin/gomin/redirect"
)

func main() {
	router := mux.NewRouter()

	environ := &env.Env{
		DB: gomin.NewRedisClient(),
	}

	redirect.Register(router, environ)

	apiRouter := router.PathPrefix("/api/v0").Subrouter()
	minify.Register(apiRouter, environ)

	log.Fatal(http.ListenAndServe(":8080", router))
}
