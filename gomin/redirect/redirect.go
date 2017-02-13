package redirect

import (
	"github.com/stupschwartz/gomin/gomin/env"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/stupschwartz/gomin/gomin/handler"
	"net/http"
)

func GetHandler(env *env.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	val, err := env.DB.Get(vars["id"]).Result()
	if err != nil {
		fmt.Fprintf(w, "Could not find id: %s", vars["id"])
	}

	http.Redirect(w, r, val, 302)

	return nil
}

func Register(router *mux.Router, environ *env.Env) {
	router.Handle("/{id}", handler.Handler{environ, GetHandler}).Methods("GET")
}