package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DenilsonNil/api-go-rest/controller/models"
	"github.com/gorilla/mux"
)

func Home(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(r, "Welcome to the Home Page!")
}

func AllPersonalidades(r http.ResponseWriter, req *http.Request) {
	fmt.Println(models.Personalidades)
	json.NewEncoder(r).Encode(models.Personalidades)
}

func FindById(r http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == key {
			json.NewEncoder(r).Encode(personalidade)
		}
	}
}
