package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DenilsonNil/api-go-rest/controller/models"
)

func Home(r http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(r, "Welcome to the Home Page!")
}

func AllPersonalidades(r http.ResponseWriter, req *http.Request) {
	fmt.Println(models.Personalidades)
	json.NewEncoder(r).Encode(models.Personalidades)
}
