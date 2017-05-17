package ghost

import (
	"log"
	"net/http"
	"os"
)

func CreateEmptyGhost() *Ghost {
	return &Ghost{}
}

// Ghost
type Ghost struct {
	middleware Middleware
}

func (g *Ghost) AddMiddleware(handler Handler) {
	g.addMiddleware(handler)
}

func (g *Ghost) AddMiddlewareByHTTPHandler(handler http.Handler) {
	g.addMiddleware(convertToGhostHandler(handler))
}

func (g *Ghost) addMiddleware(handler Handler) {
	g.middleware.Rebuild(handler)
}

func (g *Ghost) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	g.middleware.ServeHTTP(w, req)
}

func (g *Ghost) Run(port string) error {
	logger := log.New(os.Stdout, LogPrefix, 0)
	logger.Printf("StartRun. Listen port %s", port)
	err := http.ListenAndServe(port, g)
	if err != nil {
		logger.Printf("StartRun. error:%s", err.Error())
		return err
	}
	return nil
}
