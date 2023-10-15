package routes

import (
	"alpha-app/util"
	"net/http"
)

var TemplatePath = util.GetEnv("TEMPLATE_PATH", "templates/")

func RegisterRoutes() {
	http.HandleFunc("/", indexRoute)
}
