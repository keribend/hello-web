package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(Server)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "sandbox")
}
