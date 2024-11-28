package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	port := ":8080"

	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
