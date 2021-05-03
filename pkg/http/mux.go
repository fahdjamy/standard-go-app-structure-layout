package http

import (
	"context"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	"github.com/gorilla/mux"
	logger "github.com/rs/zerolog/log"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type muxRouter struct {
	apiVersion string
	address    string
}

var (
	muxDispatcher = mux.NewRouter()
)

func (m muxRouter) Serve(port string) {
	logger.Info().Msgf("Mux application server running on port: %s", port)
	muxDispatcher.PathPrefix(m.apiVersion)

	srv := &http.Server{
		Handler: muxDispatcher,
		Addr:    m.address,
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

func (m muxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (m muxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}


func (m muxRouter) NotFound(notFoundHandler func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.NotFoundHandler  = http.HandlerFunc(notFoundHandler)
}

func NewMuxRouter(apiVersion, address string) types.Router {
	return &muxRouter{apiVersion, address}
}
