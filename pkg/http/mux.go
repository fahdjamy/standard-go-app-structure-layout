package http

import (
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type muxRouter struct {}

var (
	muxDispatcher = mux.NewRouter()
)

func (m muxRouter) Serve(port string) {
	log.Info().Msgf("Mux application server running on port: %s", port)
	http.ListenAndServe(port, muxDispatcher)
}

func (m muxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (m muxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func NewMuxRouter() types.Router {
	return &muxRouter{}
}
