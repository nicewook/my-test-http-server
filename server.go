package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type myHandler struct{}

func (myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	if err := http.ListenAndServe(":"+port, http.HandlerFunc(Greeting)); err != nil {
		log.Fatal(err)
	}
}
