package ghost

import "net/http"

func convertToGhostHandler(handler http.Handler) Handler {
	return HandlerFunc(func(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		handler.ServeHTTP(w, req)
		next(w, req)
	})
}

func createEmptyMiddleware() Middleware {
	return Middleware{
		rawHandler: HandlerFunc(func(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {}),
		next:       &Middleware{},
	}
}
