package routes

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rodrigoadao/simple-go-rest-api/controllers"
)


func HandleRequest() {
    r := mux.NewRouter()
    r.HandleFunc("/", controllers.Home)
    r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
    r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("Get")
    log.Fatal(http.ListenAndServe(":8000", r))
}
