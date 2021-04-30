package myapp

import (
	"github.com/fahdjamy/standard-structure-layout/pkg/api/rest"
	http2 "github.com/fahdjamy/standard-structure-layout/pkg/http"
	"github.com/fahdjamy/standard-structure-layout/pkg/store/memory"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	"log"
	"os"
)

const port = "8080"

var (
	logger = log.New(os.Stdout, "app", log.LstdFlags|log.Lshortfile)
	service = memory.NewInMemoryDB()
)

// Router setup
func router() types.Router {
	var muxRouter types.Router = http2.NewMuxRouter()
	return muxRouter
}

func main() {
	router := router()
	rest.Routes(router, port, service, logger)
}
