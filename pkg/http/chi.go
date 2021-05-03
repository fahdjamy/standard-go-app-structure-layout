package http

import (
	"context"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	"github.com/go-chi/chi"
	logger "github.com/rs/zerolog/log"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type chiRouter struct {
	apiVersion string
	address string
}

var (
	chiDispatcher = chi.NewRouter()
)

func (c *chiRouter) Serve(port string) {
	logger.Info().Msgf("Chi application server running on port: %s", port)

	srv := &http.Server{
		Handler: muxDispatcher,
		Addr:    c.address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	ctx, ctxFunc := context.WithTimeout(context.Background(), 100*time.Second)
	<-ch
	ctxFunc()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *chiRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (c *chiRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (c *chiRouter) NotFound(custom404Handler func(w http.ResponseWriter, r *http.Request)) {
	panic("implement me")
}

func NewChiRouter(apiVersion, address string) types.Router {
	return &chiRouter{apiVersion, address}
}
