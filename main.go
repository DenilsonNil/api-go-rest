package main

import (
	"github.com/DenilsonNil/api-go-rest/controller/models"
	"github.com/DenilsonNil/api-go-rest/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Albert Einstein", Descricao: "Físico teórico conhecido por desenvolver a teoria da relatividade."},
		{Id: 2, Nome: "Marie Curie", Descricao: "Pioneira na pesquisa sobre radioatividade e primeira mulher a ganhar um Nobel."},
		{Id: 3, Nome: "Isaac Newton", Descricao: "Matemático e físico que formulou as leis do movimento e da gravitação universal."},
	}

	routes.HandleRequests()
}
