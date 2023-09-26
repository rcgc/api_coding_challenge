package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	carhandler := newCarHandler()
	http.Handle("/cars", carhandler)
	http.Handle("/cars/", carhandler)
	http.Handle("/cars/profile", carhandler)
	http.Handle("/cars/profile/", carhandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})

	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}