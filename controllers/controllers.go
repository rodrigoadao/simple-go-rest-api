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
    wr.Header().Set("Content-type", "application/json")
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

func CreateNewPersonality(wr http.ResponseWriter, r *http.Request) {
    var personality models.Personality
    json.NewDecoder(r.Body).Decode(&personality)

    database.DB.Create(&personality)
    json.NewEncoder(wr).Encode(&personality)
}

func DeletePersonality(wr http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var personality models.Personality

    database.DB.Delete(&personality, id)
    json.NewEncoder(wr).Encode(personality)
}

func EditPersonality(wr http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    var personality models.Personality
    database.DB.Find(&personality, id)

   json.NewDecoder(r.Body).Decode(&personality)
   database.DB.Save(&personality)
   json.NewEncoder(wr).Encode(&personality)

}
