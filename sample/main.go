package main

import (
	"fmt"
	"net/http"

	"github.com/mappymappy/ghost"
)

func main() {
	g := ghost.CreateEmptyGhost()
	r := http.NewServeMux()
	testfunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<div>Hello World</div>")
	}
	r.HandleFunc("/", testfunc)
	g.AddMiddleware(BasicAuthMiddleware{"user", "password"})
	g.AddMiddlewareByHTTPHandler(MyMiddleware{})
	g.AddMiddlewareByHTTPHandler(r)
	g.Run(":33333")
}

type MyMiddleware struct {
}

func (m MyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//implement yourMiddleware.
}

type BasicAuthMiddleware struct {
	authID   string
	authPass string
}

func (m BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if !m.Auth(r) {
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}
	next(w, r)
}

func (m BasicAuthMiddleware) Auth(r *http.Request) bool {
	id, pass, ok := r.BasicAuth()
	return ok && id == m.authID && pass == m.authPass
}
