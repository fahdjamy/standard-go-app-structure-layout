package http

import (
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
	"net/http"
)

type chiRouter struct {}

var (
	chiDispatcher = chi.NewRouter()
)

func (c *chiRouter) Serve(port string) {
	log.Info().Msgf("Chi application server running on port: %s", port)
	http.ListenAndServe(port, muxDispatcher)
}

func (c *chiRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (c *chiRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func NewChiRouter() types.Router {
	return &chiRouter{}
}
