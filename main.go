package main

import (
	"fmt"

	"github.com/Yurisbk/go-rest-api.git/database"
	"github.com/Yurisbk/go-rest-api.git/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleRequest()
}
