package main

import (
	"fmt"

	"github.com/rodrigoadao/simple-go-rest-api/database"
	"github.com/rodrigoadao/simple-go-rest-api/routes"
)

func main() {
    database.ConnectDatabase()
    fmt.Println("Iniciando o servidor Rest com Go")
    routes.HandleRequest()
}
