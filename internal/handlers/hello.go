package handlers


import (
	"fmt"
	"net/http"
)

func Hello(rw http.ResponseWriter, req *http.Request) {
	u := req.Context().Value("user")
	user := "unset"
	if u != nil {
		user = u.(string)
	}

	switch req.Method {
	case http.MethodGet:
		rw.Write([]byte("Hello " + user + "\n"))
	case http.MethodPost:
		rw.Write([]byte("Thanks for posting to me, " + user + "\n"))
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("invalid method")
	}
}