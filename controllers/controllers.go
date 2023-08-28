package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yurisbk/go-rest-api.git/database"
	"github.com/Yurisbk/go-rest-api.git/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidades models.Personalidade
	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&personalidades, id)
	json.NewEncoder(w).Encode(personalidades)

}
