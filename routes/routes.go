package routes

import (
	"net/http"

	"github.com/DenilsonNil/api-go-rest/controller"
)

func HandleRequests() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/api/personalidades", controller.AllPersonalidades)
	http.ListenAndServe(":8080", nil)
}
