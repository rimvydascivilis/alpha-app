package routes

import (
	"log"
	"net/http"
)

func HealthCheckRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("{\"message\": \"OK\"}"))
	if err != nil {
		log.Println("Error writing response: ", err)
	}
}
