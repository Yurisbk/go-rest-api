package main

import (
	"fmt"

	"github.com/Yurisbk/go-rest-api.git/models"
	"github.com/Yurisbk/go-rest-api.git/routes"
)

func main() {
	models.Personalidades = []models.Pesonalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando seridor Rest com Go")
	routes.HandleRequest()
}
