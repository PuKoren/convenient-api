package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

    "github.com/PuKoren/convenient-api/routes"
    "github.com/PuKoren/convenient-api/models"
)

func main() {
    models.Init()

    router := mux.NewRouter()

    routes.RegisterHandlersUser(router)

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(router)))
}
