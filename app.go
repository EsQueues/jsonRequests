package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start http server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
