package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hello() string {
	return "Hello, welcome to gopher"
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(Hello()))
	}))

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	fmt.Println("Server listening on http://localhost:3000")
	log.Fatal(server.ListenAndServe())
}
