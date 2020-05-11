package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type helloServer struct {
	portNum int
}

func New(portNum int) *helloServer {
	return &helloServer{
		portNum: portNum, 
	}
}

func (h *helloServer) sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!!")
}

func (h *helloServer) HttpServer() *http.Server {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", h.sayHello).Methods("GET")
	return &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", h.portNum),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
