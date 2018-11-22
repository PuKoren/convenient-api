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

    allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    allowedOrigins := handlers.AllowedOrigins([]string{"*"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

    log.Println(http.ListenAndServe(":8000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))

    models.Close()
}
