package routes

import (
	"net/http"

	"github.com/DenilsonNil/api-go-rest/controller"
	"github.com/DenilsonNil/api-go-rest/controller/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.SetContentTypeJSON)
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/personalidades", controller.AllPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades", controller.Create).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controller.FindById).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controller.DeleteById).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controller.EditById).Methods("Put")
	http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
}
