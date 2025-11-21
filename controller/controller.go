package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DenilsonNil/api-go-rest/database"
	"github.com/DenilsonNil/api-go-rest/models"
	"github.com/gorilla/mux"
)

func Home(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(r, "Welcome to the Home Page!")
}

func Create(response http.ResponseWriter, request *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(request.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)
	json.NewEncoder(response).Encode(personalidade)
}

func AllPersonalidades(r http.ResponseWriter, req *http.Request) {
	r.Header().Set("Content-Type", "application/json")
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(r).Encode(p)
}

func FindById(r http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	// for _, personalidade := range models.Personalidades {
	// 	if strconv.Itoa(personalidade.Id) == key {
	// 		json.NewEncoder(r).Encode(personalidade)
	// 	}
	// }

	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(r).Encode(p)
}

func DeleteById(r http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var p models.Personalidade
	database.DB.Delete(&p, id)
	json.NewEncoder(r).Encode(p)
}

func EditById(r http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewDecoder(req.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(r).Encode(personalidade)
}
