package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"
func main() {
	fmt.Println("Starting server on port", port)
	srv := &http.Server{
		Addr:    port,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}