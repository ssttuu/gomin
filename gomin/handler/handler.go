package handler

import (
	"gopkg.in/redis.v5"
	"log"
	"net/http"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Env struct {
	DB *redis.Client
	//Port string
	//Host string
}

type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}
