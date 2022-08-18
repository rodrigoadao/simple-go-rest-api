package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rodrigoadao/simple-go-rest-api/controllers"
	"github.com/rodrigoadao/simple-go-rest-api/middleware"
)


func HandleRequest() {
    r := mux.NewRouter()
    r.Use(middleware.ContentTypeMiddleware)
    r.HandleFunc("/", controllers.Home)
    r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
    r.HandleFunc("/api/personalities", controllers.CreateNewPersonality).Methods("Post")
    r.HandleFunc("/api/personalities/{id}", controllers.GetPersonality).Methods("Get")
    r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
    r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
