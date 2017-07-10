ghost [![GoDoc](http://godoc.org/github.com/mappymappy/ghost?status.svg)](http://godoc.org/github.com/mappymappy/ghost)
======

![ghost](https://cloud.githubusercontent.com/assets/6446570/26194856/cb108864-3bf4-11e7-9915-24d019f39a64.png)

## Introduction

Ghost is micro framework for web application written by golang.

The design concepts of Ghost is tiny,thin,and confirm to `net/http`.

Most framework provide various feature,such as Logger,Mux,Template,etc.

Ghost does not have Them.

Instead only the middleware chain mechanism and only the minimal procedures

using `net/http` required for web applications

possible to register that satisfy `net/http.HandleFunc` interface as Middleware.

Also, your middleware does not need to depend on Ghost, just depend on `net/http`

## install

```
  go get github.com/mappymappy/ghost
```

## tutorial

please read `sample/main.go`

```
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

```

## How to add Middleware?

ghost have two ways for register middleware.

### ghost.AddMiddlewareByHttpHandler

Middleware registered by this method can not interrupt the chain.

If your middleware does not need it,you better use this

### ghost.AddMiddleware

Please use when your middleware need the right to interrupt the chain.

```
func (m BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// authentication failure
	if !m.Auth(r) {
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return
	}
	next(w, r)
}
```

## Optional Middleware

* [panic_recover](https://github.com/mappymappy/panic_recover)
* [http_logger](https://github.com/mappymappy/http_logger)

## Author
[marnie_ms4](https://github.com/mappymappy?tab=repositories)
