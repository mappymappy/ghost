package ghost

import "net/http"

type Middleware struct {
	rawHandler Handler
	next       *Middleware
	hasNext    bool
}

func (m *Middleware) Rebuild(handler Handler) {
	var last *Middleware
	current := m
	nullMiddleware := createEmptyMiddleware()
	//at first rebuild
	if current.rawHandler == nil {
		m.rawHandler = handler
		m.next = &nullMiddleware
		return
	}
	for {
		if current == nil {
			return
		}
		if !current.hasNext {
			last = current
			break
		}
		current = current.next
	}
	last.next = &Middleware{
		rawHandler: handler,
		next:       &nullMiddleware,
	}
	last.hasNext = true
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.rawHandler.ServeHTTP(w, req, m.next.ServeHTTP)
}
