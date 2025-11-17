package routes

import (
	"net/http"

	"github.com/DenilsonNil/api-go-rest/controller"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/personalidades", controller.AllPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controller.FindById).Methods("Get")
	http.ListenAndServe(":8080", r)
}
