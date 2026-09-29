package main

import (
	"net"
	"net/http"
	"os"
)

func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("GET /health", health)
	if err := http.Serve(listener, nil); err != nil {
		panic(err)
	}
}
