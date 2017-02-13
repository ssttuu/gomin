package minify

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stupschwartz/gomin/gomin/env"
	"github.com/stupschwartz/gomin/gomin/handler"
	"math/rand"
	"net/http"
	"encoding/json"
)

type MinifyBody struct {
	Url string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetAllHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	val, err := env.DB.Keys("*").Result()
	if err != nil {
		fmt.Fprint(w, "Could not find any keys")
	}
	fmt.Fprintf(w, "%s", val)
	return nil
}

func GetHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	val, err := env.DB.Get(vars["id"]).Result()
	if err != nil {
		fmt.Fprintf(w, "Could not find id: %s", vars["id"])
	}

	fmt.Fprint(w, val)

	return nil
}

func PostHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	var body MinifyBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {

	}
	random_id := RandStringBytes(8)

	redis_err := env.DB.Set(random_id, body.Url, 0).Err()
	if redis_err != nil {
		// Handle error
	}

	fmt.Fprintf(w, "Random_id: %s, Url: %s", random_id, body.Url)

	return nil
}

func PutHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func DeleteHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	s := router.PathPrefix("/minify").Subrouter()
	s.Handle("/", handler.Handler{environ, GetAllHandler}).Methods("GET")
	s.Handle("/{id}", handler.Handler{environ, GetHandler}).Methods("GET")
	s.Handle("/", handler.Handler{environ, PostHandler}).Methods("POST")
	s.Handle("/{id}", handler.Handler{environ, PutHandler}).Methods("PUT")
	s.Handle("/{id}", handler.Handler{environ, DeleteHandler}).Methods("DELETE")
}
