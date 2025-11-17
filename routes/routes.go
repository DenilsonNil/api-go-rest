package routes

import (
	"net/http"

	"github.com/DenilsonNil/api-go-rest/controller"
)

func HandleRequests() {
	http.HandleFunc("/", controller.Home)
	http.ListenAndServe(":8080", nil)
}
