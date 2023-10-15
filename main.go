package main

import (
	"alpha-app/routes"
	"alpha-app/util"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes.RegisterRoutes()

	port := util.GetEnv("PORT", "8080")

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
