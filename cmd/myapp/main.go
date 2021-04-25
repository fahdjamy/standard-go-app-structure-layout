package myapp

import (
	"github.com/fahdjamy/standard-structure-layout/pkg/api/rest"
	http2 "github.com/fahdjamy/standard-structure-layout/pkg/http"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
)

const port = "8080"

// Router setup
func router() types.Router {
	var muxRouter types.Router = http2.NewMuxRouter()
	return muxRouter
}

func main() {
	router := router()
	rest.Routes(router, port)
}
