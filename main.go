package main

import (
	"fmt"
	"net/http"
	"time"
	"context"
)


type middleware struct {
	mux http.Handler
}

func (m middleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := context.WithValue(req.Context(), "user", "unknown")
	ctx = context.WithValue(ctx, "__requestStartTimer__", time.Now())
	req = req.WithContext(ctx)

	m.mux.ServeHTTP(rw, req)

	start := req.Context().Value("__requestStartTimer__").(time.Time)
	fmt.Println("request duration: ", time.Now().Sub(start))
}

func hello(rw http.ResponseWriter, req *http.Request){
	// if _,err := rw.Write([]byte("hello")); err != nil{
	// 	fmt.Print("error when writing response for /user request")
	// 	rw.WriteHeader(http.StatusInternalServerError)	
	// }		
	switch req.Method {
case http.MethodGet:
  if _, err := rw.Write([]byte("Hello World\n")); err != nil {
    fmt.Println("error when writing response for /hello request")
    rw.WriteHeader(http.StatusInternalServerError)
  }
case http.MethodPost:
  if _, err := rw.Write([]byte("Thanks for posting to me\n")); err != nil {
    fmt.Println("error when writing response for /hello request")
    rw.WriteHeader(http.StatusInternalServerError)
  }
}
}


func helloTwo(rw http.ResponseWriter, req *http.Request){
	// if _,err := rw.Write([]byte("hello")); err != nil{
	// 	fmt.Print("error when writing response for /user request")
	// 	rw.WriteHeader(http.StatusInternalServerError)	
	// }		
	switch req.Method {
case http.MethodGet:
  if _, err := rw.Write([]byte("Hello World\n")); err != nil {
    fmt.Println("error when writing response for /hello request")
    rw.WriteHeader(http.StatusInternalServerError)
  }
case http.MethodPost:
  if _, err := rw.Write([]byte("Thanks for posting to me\n")); err != nil {
    fmt.Println("error when writing response for /hello request")
    rw.WriteHeader(http.StatusInternalServerError)
  }
}
}


func main() {
	fmt.Println("start server")

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/user", helloTwo)

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      middleware{mux},
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}