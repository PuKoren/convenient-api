package main

import (
    _ "encoding/json"
    "net/http"
    "log"

    "github.com/gorilla/mux"

    _ "github.com/PuKoren/convenient-api/routes/form"
)

func main() {
    router := mux.NewRouter()



    log.Fatal(http.ListenAndServe(":8000", router))
}
