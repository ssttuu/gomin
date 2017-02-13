package minify

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stupschwartz/gomin/gomin/handler"
	"net/http"
)

func GetAllHandler(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprint(w, "Hi there, I love everything!")
	return nil
}

func GetHandler(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hi there, I love %s!", vars["id"])
	return nil
}

func PostHandler(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func PutHandler(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func DeleteHandler(env *handler.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Register(router *mux.Router, env *handler.Env) {
	s := router.PathPrefix("/minify").Subrouter()
	s.Handle("/", handler.Handler{env, GetAllHandler}).Methods("GET")
	s.Handle("/{id}", handler.Handler{env, GetHandler}).Methods("GET")
	s.Handle("/", handler.Handler{env, PostHandler}).Methods("POST")
	s.Handle("/{id}", handler.Handler{env, PutHandler}).Methods("PUT")
	s.Handle("/{id}", handler.Handler{env, DeleteHandler}).Methods("DELETE")
}
