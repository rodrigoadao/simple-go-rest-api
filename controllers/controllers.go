package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigoadao/simple-go-rest-api/database"
	"github.com/rodrigoadao/simple-go-rest-api/models"
)

func Home (wr http.ResponseWriter, r *http.Request) {
    fmt.Fprint(wr, "Home Page")
}

func AllPersonalities(wr http.ResponseWriter, r *http.Request) {
    var personalities []models.Personality
    database.DB.Find(&personalities)

    json.NewEncoder(wr).Encode(personalities)
}

func GetPersonality(wr http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var personality models.Personality
    database.DB.Find(&personality, id)

    json.NewEncoder(wr).Encode(personality)
}
