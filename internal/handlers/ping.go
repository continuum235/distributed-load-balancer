package handlers

import "net/http"

func Ping(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		rw.Write([]byte("pong\n"))
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}