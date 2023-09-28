// @title Cars Restful API with Swagger
// @version 0.1
// @description Simple swagger implementation in Go HTTP
// @termsOfService  http://swagger.io/terms/
//
// @contact.name Roberto Guzmán
// @contact.email roberto140298@gmail.com
//
// @license.name  Apache 2.0
// @license.url   https://opensource.org/license/mit/
//
// @host localhost:8080
// @BasePath /
//
// @securityDefinitions.basic  BasicAuth
//
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "example/cars/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func main() {
	port := ":8080"

	carhandler := newCarHandler()
	http.Handle("/cars", carhandler)
	http.Handle("/cars/", carhandler)

	http.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		httpSwagger.WrapHandler(w, r)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})

	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}