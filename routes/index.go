package routes

import (
	"alpha-app/util"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Data struct {
	Color   string
	Message string
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for URL: ", r.URL)

	tmp := template.Must(template.ParseFiles(filepath.Join(TemplatePath, "index.html")))

	name := "World"
	if n := r.URL.Query().Get("name"); n != "" {
		name = n
	}

	color := util.GetEnv("COLOR", "blue")
	message := fmt.Sprintf("Hello, %s!", name)
	data := Data{
		Color:   string(color),
		Message: string(message),
	}

	log.Println("Sending response: ", data)
	err := tmp.Execute(w, data)
	if err != nil {
		log.Println("Error executing template: ", err)
	}
}
