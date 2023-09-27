//	Cars API
//
//	Api to administrate cars requests
//
//		Schemes: http
//
//   	Host: localhost:8080
//  	BasePath: /
//
//	Version: 1.0
//	License: MIT http://opensource.org/licenses/MIT
//
//      Consumes:
//   - application/json
//     Produces:
//   - application/json
//
//   swagger:meta
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})

	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}