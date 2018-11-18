package form

import (
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"

    "github.com/PuKoren/convenient-api/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")

    json.NewEncoder(w).Encode(models.User{})
}

func RegisterHandlers(router *mux.Router) {

    router.HandleFunc("/form/user/v1", GetUser).Methods("GET")
}
