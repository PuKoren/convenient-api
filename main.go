package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"

    "github.com/PuKoren/convenient-api/routes/form"
)

func main() {

    router := mux.NewRouter()

    form.RegisterHandlers(router)

    log.Fatal(http.ListenAndServe(":8000", router))
}
