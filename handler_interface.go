package ghost

import "net/http"

type Handler interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc)
}
