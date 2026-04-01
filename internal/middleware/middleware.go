package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Middleware struct {
	Mux http.Handler
}

func (m Middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), "user", "unknown")
	ctx = context.WithValue(ctx, "__requestStartTimer__", time.Now())

	req = req.WithContext(ctx)

	m.Mux.ServeHTTP(rw, req)

	start := req.Context().Value("__requestStartTimer__").(time.Time)
	fmt.Println("request duration:", time.Since(start))
}