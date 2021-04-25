package types

import "net/http"

type Error struct {
	Message string `json:"message"`
}

type Router interface {
	Serve (port string)
	Get (uri string, f func(w http.ResponseWriter, r *http.Request))
	Post (uri string, f func(w http.ResponseWriter, r *http.Request))
}
