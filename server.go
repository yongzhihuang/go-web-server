package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	port := ":1337"
	log.Println("Listening on port %s", port)
	http.ListenAndServe(port, nil)
}
