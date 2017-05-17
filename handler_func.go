package ghost

import "net/http"

type HandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	h(w, req, next)
}
